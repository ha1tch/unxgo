package vm

import (
	"fmt"
	"os"
)

// IO handles input/output operations for the Uxn VM

// EmuDei reads from a device port
func (u *Uxn) EmuDei(addr uint8) uint8 {
	return u.Dev[addr]
}

// EmuDeo writes to a device port
func (u *Uxn) EmuDeo(addr uint8, value uint8) {
	u.Dev[addr] = value
	switch addr {
	case 0x18:
		fmt.Print(string(u.Dev[addr]))
	case 0x19:
		fmt.Fprint(os.Stderr, string(u.Dev[addr]))
	}
}

// ConsoleInput handles console input and triggers VM evaluation
func (u *Uxn) ConsoleInput(c int, inputType int) bool {
	if c == -1 {
		c = 0
		inputType = 4
	}
	u.Dev[0x12] = uint8(c)
	u.Dev[0x17] = uint8(inputType)
	return u.Eval(uint16(u.Dev[0x10])<<8|uint16(u.Dev[0x11])) && inputType != 4
}