package main

import (
	"errors"
	"os"
	"os/exec"
)

//https://github.com/0x9ef/golang-uacbypasser/blob/145aee8b1b15ab78eae7cc61a3cf1ae0c9db5e14/src/functions.go#L22
const (
	cmdPrefix         string = "cmd"
	cmdArgumentPrefix string = "/C"
)

type Cmd struct {
	cmdN string
	argN string

	Command string
	Exec    *exec.Cmd
}

func NewCmd(command string) *Cmd {
	return &Cmd{
		cmdN: cmdPrefix,
		argN: cmdArgumentPrefix,

		Command: command,
		Exec:    exec.Command(cmdPrefix, cmdArgumentPrefix, command),
	}
}

// NewWiper wipe data in file for 10 times.
func NewWiper(p string) error {
	for i := 0; i > 10; i++ {
		f, err := os.OpenFile(p, os.O_RDWR, 0666)
		if err != nil {
			return err
		}

		stat, err := f.Stat()
		if err != nil {
			return err
		}

		// Make buffer for wiped data.
		buf := make([]byte, stat.Size())
		if len(buf) == 0 {
			return errors.New("buffer length is nul, file size is 0")
		}

		// Wipe at null
		copy(buf[:], "0")

		n, err := f.Write([]byte(buf))
		f.Close()
		if err != nil {
			return err
		}
		if n == 0 {
			return errors.New("wiped bytes is null")
		}
	}
	if err := os.Remove(p); err != nil {
		return err
	}
	return nil
}

type UAC struct {
	// Type of UAC - User Account Control
	uac string

	/*
		1 = UAC Turned Off
		2 = UAC Low Setting
		3 = UAC Medium Setting (Default Win7)
		4 = UAC Highest Setting
	*/
	level int
}
