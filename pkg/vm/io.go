package vm

import (
	"fmt"
	"os"
)

// IO handles input/output operations for the Uxn VM

// EmuDei reads from a device port
func (u *Uxn) EmuDei(addr uint8) uint8 {
	return u.dev[addr]
}

// EmuDeo writes to a device port
func (u *Uxn) EmuDeo(addr uint8, value uint8) {
	u.dev[addr] = value
	switch addr {
	case 0x18:
		fmt.Print(string(u.dev[addr]))
	case 0x19:
		fmt.Fprint(os.Stderr, string(u.dev[addr]))
	}
}

// ConsoleInput handles console input and triggers VM evaluation
func (u *Uxn) ConsoleInput(c int, inputType int) bool {
	if c == -1 {
		c = 0
		inputType = 4
	}
	u.dev[0x12] = uint8(c)
	u.dev[0x17] = uint8(inputType)
	return u.Eval(uint16(u.dev[0x10])<<8|uint16(u.dev[0x11])) && inputType != 4
}