package commandline

import (
	"unsafe"

	"golang.org/x/sys/windows"
)

var kernel32 = windows.NewLazySystemDLL("kernel32")

var procGetCommandLine = kernel32.NewProc("GetCommandLineW")

func Get() string {
	lpwstr, _, _ := procGetCommandLine.Call()
	utf8 := windows.UTF16PtrToString((*uint16)(unsafe.Pointer(lpwstr)))
	return utf8
}
