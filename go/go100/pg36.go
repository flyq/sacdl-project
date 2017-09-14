package main

import (
	"fmt"
	"syscall"
)

type (
	HANDLE uintptr
	WORD   uint16
	DWORD  uint32
)

const (
	STD_OUTPUT_HANDLE    = 0xFFFFFFF5
	FOREGROUND_BLUE      = 0x01
	FOREGROUND_GREEN     = 0x02
	FOREGROUND_RED       = 0x04
	FOREGROUND_INTENSITY = 0x08
	BACKGROUND_BLUE      = 0x10
	BACKGROUND_GREEN     = 0x20
	BACKGROUND_RED       = 0x40
	BACKGROUND_INTENSITY = 0x80
)

var (
	modkernel32                 = syscall.NewLazyDLL("kernel32.dll")
	procGetStdHandle            = modkernel32.NewProc("GetStdHandle")
	procSetConsoleTextAttribute = modkernel32.NewProc("SetConsoleTextAttribute")
)

func GetStdHandle(nStdHandle DWORD) HANDLE {
	ret, _, _ := procGetStdHandle.Call(uintptr(nStdHandle))
	return HANDLE(ret)
}
func SetConsoleTextAttribute(hConsoleOutput HANDLE, wAttributes WORD) bool {
	ret, _, _ := procSetConsoleTextAttribute.Call(
		uintptr(hConsoleOutput),
		uintptr(wAttributes))
	return ret != 0
}
func main() {
	for color := 1; color < 16; color++ {
		textcolor(color) /*设置文本颜色*/
		fmt.Printf("This is color %d\n", color)
		textcolor(128 + color)
		fmt.Printf("This is blinking\n")
	}
}
func textcolor(color int) {
	hOut := GetStdHandle(STD_OUTPUT_HANDLE)
	SetConsoleTextAttribute(hOut, WORD(color))
}
