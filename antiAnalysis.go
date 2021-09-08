package main

import (
	"syscall"
	"unsafe"
)

func MessageBox(hwnd uintptr, caption, title string, flags uint) int {
	ret, _, _ := syscall.NewLazyDLL("user32.dll").NewProc("MessageBoxW").Call(
		uintptr(hwnd),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(caption))),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(title))),
		uintptr(flags))

	return int(ret)
}

// MessageBoxPlain of Win32 API.
func MessageBoxPlain(title, caption string) int {
	const (
		NULL  = 0
		MB_OK = 0
	)
	return MessageBox(NULL, caption, title, MB_OK)
}

// string hashing

func stringHashing(target string) int64 {
	buffer := []byte(target)
	var result int64 = 0

	for each := range buffer {
		result = ((result + int64(each)) % 0x69696969) ^ int64(each)
	}
	return result
}
