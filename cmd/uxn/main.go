package main

import (
	"fmt"
	"os"

	"github.com/yourusername/uxngo/pkg/vm"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("usage: %s file.rom [args..]\n", os.Args[0])
		returns
	}

	machine := vm.NewUxn()

	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "uxnmin: Error %s\n", os.Args[1])
		return
	}
	defer f.Close()

	if len(os.Args) > 2 {
		machine.Dev[0x17] = 1
	}

	// Read ROM file
	f.Read(machine.Ram[0x0100:])

	if machine.Eval(0x0100) && machine.Dev[0x10] != 0 {
		// Process arguments
		for i := 2; i < len(os.Args); i++ {
			for _, c := range os.Args[i] {
				machine.ConsoleInput(int(c), 0x2)
			}
			inputType := 0x3
			if i == len(os.Args)-1 {
				inputType = 0x4
			}
			machine.ConsoleInput(int('\n'), inputType)
		}

		// Process stdin
		for machine.Dev[0x0f] == 0 {
			c := make([]byte, 1)
			_, err := os.Stdin.Read(c)
			if err != nil {
				break
			}
			if !machine.ConsoleInput(int(c[0]), 0x1) {
				break
			}
		}
	}

	os.Exit(int(machine.Dev[0x0f] & 0x7f))
}