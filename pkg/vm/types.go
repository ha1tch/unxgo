// Package vm implements the core Uxn virtual machine
package vm

// Stack represents a Uxn stack with its data and pointer
type Stack struct {
	dat [0x100]uint8
	ptr uint8
}

// Uxn represents the complete state of the Uxn virtual machine
type Uxn struct {
	Ram [0x10000]uint8 // Ram contains the VM's memory
	Dev [0x100]uint8   // Dev contains device I/O ports
	wst Stack          // Working stack
	rst Stack          // Return stack
}

// Stack operations
func (s *Stack) inc() uint8 {
	val := s.dat[s.ptr]
	s.ptr++
	return val
}

func (s *Stack) dec() uint8 {
	s.ptr--
	return s.dat[s.ptr]
}