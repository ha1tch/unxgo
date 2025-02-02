# UxnGo

UxnGo is a Go implementation of the Uxn virtual machine. Uxn is a minimalist stack-based virtual machine designed for small, portable computing systems.

## Overview

The Uxn virtual machine is a stack-based computer with the following characteristics:

- 64KB addressable memory space
- Two stacks (working stack and return stack)
- 256 bytes of device memory for I/O operations
- 32 core instructions with multiple addressing modes
- 8-bit and 16-bit operation modes

This implementation provides a complete Uxn emulator in Go, capable of running Uxn ROMs and interfacing with standard input/output devices.

## Features

- Complete implementation of all Uxn instructions
- Stack operations (working and return stacks)
- Device I/O support
- Console input/output handling
- ROM loading and execution
- Command-line argument support

## Project Structure

```
uxngo/
├── cmd/
│   └── uxn/
│       └── main.go         # Main executable
├── pkg/
│   └── vm/
│       ├── types.go        # Core type definitions
│       ├── io.go           # I/O operations
│       ├── ops.go          # Instruction implementations
│       └── vm.go           # VM core and evaluation
└── go.mod                  # Module definition
```

## Installation

1. Ensure you have Go 1.21 or later installed
2. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/uxngo.git
   ```
3. Build the project:
   ```bash
   cd uxngo
   go build ./cmd/uxn
   ```

## Usage

Run a ROM file:
```bash
./uxn path/to/rom.rom
```

Run a ROM file with arguments:
```bash
./uxn path/to/rom.rom arg1 arg2
```

## Instruction Set

The Uxn VM supports the following core instructions:

### Stack Operations
- `BRK`: Break execution
- `INC`: Increment
- `POP`: Remove value
- `NIP`: Remove value below
- `SWP`: Swap two values
- `ROT`: Rotate three values
- `DUP`: Duplicate value
- `OVR`: Duplicate value below

### Flow Control
- `JCI`: Jump if condition
- `JMI`: Jump immediate
- `JSI`: Jump and store
- `JMP`: Jump
- `JCN`: Jump conditional
- `JSR`: Jump subroutine

### Memory Operations
- `LIT`: Load literal
- `LDZ`: Load from zero page
- `STZ`: Store to zero page
- `LDR`: Load relative
- `STR`: Store relative
- `LDA`: Load absolute
- `STA`: Store absolute

### Arithmetic and Logic
- `ADD`: Addition
- `SUB`: Subtraction
- `MUL`: Multiplication
- `DIV`: Division
- `AND`: Logical AND
- `ORA`: Logical OR
- `EOR`: Logical XOR
- `SFT`: Bit shift

### Device I/O
- `DEI`: Device in
- `DEO`: Device out

## Device Support

The implementation includes basic device support:

- `0x18`: Standard output (stdout)
- `0x19`: Error output (stderr)
- `0x12`: Console input character
- `0x17`: Console input type

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request. For major changes, please open an issue first to discuss what you would like to change.

## License

[Your chosen license]

## Acknowledgments

This implementation is based on the original Uxn specification and C implementation by [Hundred Rabbits](https://100r.co/site/uxn.html).

## Resource Links

- [Uxn Specification](https://wiki.xxiivv.com/site/uxn.html)
- [Original C Implementation](https://git.sr.ht/~rabbits/uxn)