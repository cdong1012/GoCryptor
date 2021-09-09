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
func bufferHashing(target []byte) uint32 {
	var result uint32 = 0

	for each := range target {
		result = uint32(each) + ror(result, 13)
	}
	return result
}

func ror(input uint32, bits uint32) uint32 {
	return (input >> bits) | (input << (32 - bits))
}
