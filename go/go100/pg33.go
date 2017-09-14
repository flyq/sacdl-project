package main

import (
	"fmt"
	"syscall"
	"unsafe"
)

type (
	HANDLE uintptr
	WORD   uint16
	DWORD  uint32
	TCHAR  rune
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

type COORD struct {
	X, Y int16
}
type SMALL_RECT struct {
	Left, Top, Right, Bottom int16
}
type CONSOLE_SCREEN_BUFFER_INFO struct {
	DwSize              COORD
	DwCursorPosition    COORD
	WAttributes         WORD
	SrWindow            SMALL_RECT
	DwMaximumWindowSize COORD
}

var (
	modkernel32                    = syscall.NewLazyDLL("kernel32.dll")
	procFillConsoleOutputAttribute = modkernel32.NewProc("FillConsoleOutputAttribute")
	procFillConsoleOutputCharacter = modkernel32.NewProc("FillConsoleOutputCharacterW")
	procGetStdHandle               = modkernel32.NewProc("GetStdHandle")
	procGetConsoleScreenBufferInfo = modkernel32.NewProc("GetConsoleScreenBufferInfo")
	procSetConsoleCursorPosition   = modkernel32.NewProc("SetConsoleCursorPosition")
	procSetConsoleTextAttribute    = modkernel32.NewProc("SetConsoleTextAttribute")
)

func COORD2DWORD(c COORD) DWORD {
	return DWORD(int32(c.Y)<<16 + int32(c.X))
}
func FillConsoleOutputAttribute(hConsoleOutput HANDLE, wAttribute WORD, nLength DWORD, dwWriteCoord COORD) *DWORD {
	var lpNumberOfAttrsWritten DWORD
	ret, _, _ := procFillConsoleOutputAttribute.Call(
		uintptr(hConsoleOutput),
		uintptr(wAttribute),
		uintptr(nLength),
		uintptr(COORD2DWORD(dwWriteCoord)),
		uintptr(unsafe.Pointer(&lpNumberOfAttrsWritten)))
	if ret == 0 {
		return nil
	}
	return &lpNumberOfAttrsWritten
}
func FillConsoleOutputCharacter(hConsoleOutput HANDLE, cCharacter TCHAR, nLength DWORD, dwWriteCoord COORD) *DWORD {
	var lpNumberOfAttrsWritten DWORD
	ret, _, _ := procFillConsoleOutputCharacter.Call(
		uintptr(hConsoleOutput),
		uintptr(cCharacter),
		uintptr(nLength),
		uintptr(COORD2DWORD(dwWriteCoord)),
		uintptr(unsafe.Pointer(&lpNumberOfAttrsWritten)))
	if ret == 0 {
		return nil
	}
	return &lpNumberOfAttrsWritten
}
func GetStdHandle(nStdHandle DWORD) HANDLE {
	ret, _, _ := procGetStdHandle.Call(uintptr(nStdHandle))
	return HANDLE(ret)
}
func GetConsoleScreenBufferInfo(hConsoleOutput HANDLE) *CONSOLE_SCREEN_BUFFER_INFO {
	var csbi CONSOLE_SCREEN_BUFFER_INFO
	ret, _, _ := procGetConsoleScreenBufferInfo.Call(
		uintptr(hConsoleOutput),
		uintptr(unsafe.Pointer(&csbi)))
	if ret == 0 {
		return nil
	}
	return &csbi
}
func SetConsoleCursorPosition(hConsoleOutput HANDLE, dwCursorPosition COORD) bool {
	ret, _, _ := procSetConsoleCursorPosition.Call(
		uintptr(hConsoleOutput),
		uintptr(COORD2DWORD(dwCursorPosition)))
	return ret != 0
}
func SetConsoleTextAttribute(hConsoleOutput HANDLE, wAttributes WORD) bool {
	ret, _, _ := procSetConsoleTextAttribute.Call(
		uintptr(hConsoleOutput),
		uintptr(wAttributes))
	return ret != 0
}
func clrscr() {
	hConsole := GetStdHandle(STD_OUTPUT_HANDLE)
	coordScreen := COORD{0, 0}
	var dwConSize DWORD
	var csbi *CONSOLE_SCREEN_BUFFER_INFO
	csbi = GetConsoleScreenBufferInfo(hConsole)
	dwConSize = DWORD(csbi.DwSize.X * csbi.DwSize.Y)
	FillConsoleOutputCharacter(hConsole, TCHAR(' '), dwConSize, coordScreen)
	csbi = GetConsoleScreenBufferInfo(hConsole)
	FillConsoleOutputAttribute(hConsole, csbi.WAttributes, dwConSize, coordScreen)
	SetConsoleCursorPosition(hConsole, coordScreen)
}
func gotoxy(x, y int16) {
	loc := COORD{x, y}
	hConsole := GetStdHandle(STD_OUTPUT_HANDLE)
	SetConsoleCursorPosition(hConsole, loc)
}
func textbackground(color int) {
	hOut := GetStdHandle(STD_OUTPUT_HANDLE)
	SetConsoleTextAttribute(hOut, WORD(color))
}
func main() {
	clrscr() /*清屏函数*/
	textbackground(2)
	gotoxy(1, 5) /*定位函数*/
	fmt.Printf("Output at row 5 column 1\n")
	textbackground(3)
	gotoxy(20, 10)
	fmt.Printf("Output at row 10 column 20\n")
}
