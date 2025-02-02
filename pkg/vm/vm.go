package vm

// NewUxn creates a new Uxn virtual machine instance
func NewUxn() *Uxn {
	return &Uxn{}
}

// Eval executes the Uxn virtual machine starting from the given program counter
func (u *Uxn) Eval(pc uint16) bool {
	if pc == 0 || u.dev[0x0f] != 0 {
		return false
	}

	step := uint32(0x80000)
	for step > 0 {
		step--
		
		instr := u.ram[pc]
		pc++

		switch instr {
		case 0x00: // BRK
			return true

		case 0x20: // JCI
			if u.wst.dec() != 0 {
				a := uint16(u.ram[pc])<<8 | uint16(u.ram[pc+1])
				pc += a + 2
			} else {
				pc += 2
			}

		case 0x40: // JMI
			a := uint16(u.ram[pc])<<8 | uint16(u.ram[pc+1])
			pc += a + 2

		case 0x60: // JSI
			c := pc + 2
			u.rst.dat[u.rst.ptr] = uint8(c >> 8)
			u.rst.ptr++
			u.rst.dat[u.rst.ptr] = uint8(c)
			u.rst.ptr++
			a := uint16(u.ram[pc])<<8 | uint16(u.ram[pc+1])
			pc += a + 2

		case 0x80: // LIT
			u.wst.dat[u.wst.ptr] = u.ram[pc]
			u.wst.ptr++
			pc++

		case 0xa0: // LI2
			u.wst.dat[u.wst.ptr] = u.ram[pc]
			u.wst.ptr++
			pc++
			u.wst.dat[u.wst.ptr] = u.ram[pc]
			u.wst.ptr++
			pc++

		case 0xc0: // LIr
			u.rst.dat[u.rst.ptr] = u.ram[pc]
			u.rst.ptr++
			pc++

		case 0xe0: // L2r
			u.rst.dat[u.rst.ptr] = u.ram[pc]
			u.rst.ptr++
			pc++
			u.rst.dat[u.rst.ptr] = u.ram[pc]
			u.rst.ptr++
			pc++

		default:
			// Handle other opcodes
			mode := instr & 0xe0
			opcode := instr & 0x1f

			switch opcode {
			case 0x01: // INC
				u.handleInc(mode)
			case 0x02: // POP
				u.handlePop(mode)
			case 0x03: // NIP
				u.handleNip(mode)
			case 0x04: // SWP
				u.handleSwp(mode)
			case 0x05: // ROT
				u.handleRot(mode)
			case 0x06: // DUP
				u.handleDup(mode)
			case 0x07: // OVR
				u.handleOvr(mode)
			case 0x08: // EQU
				u.handleEqu(mode)
			case 0x09: // NEQ
				u.handleNeq(mode)
			case 0x0a: // GTH
				u.handleGth(mode)
			case 0x0b: // LTH
				u.handleLth(mode)
			case 0x0c: // JMP
				pc = u.handleJmp(mode, pc)
			case 0x0d: // JCN
				pc = u.handleJcn(mode, pc)
			case 0x0e: // JSR
				pc = u.handleJsr(mode, pc)
			case 0x0f: // STH
				u.handleSth(mode)
			case 0x10: // LDZ
				u.handleLdz(mode)
			case 0x11: // STZ
				u.handleStz(mode)
			case 0x12: // LDR
				u.handleLdr(mode, pc)
			case 0x13: // STR
				u.handleStr(mode, pc)
			case 0x14: // LDA
				u.handleLda(mode)
			case 0x15: // STA
				u.handleSta(mode)
			case 0x16: // DEI
				u.handleDei(mode)
			case 0x17: // DEO
				u.handleDeo(mode)
			case 0x18: // ADD
				u.handleAdd(mode)
			case 0x19: // SUB
				u.handleSub(mode)
			case 0x1a: // MUL
				u.handleMul(mode)
			case 0x1b: // DIV
				u.handleDiv(mode)
			case 0x1c: // AND
				u.handleAnd(mode)
			case 0x1d: // ORA
				u.handleOra(mode)
			case 0x1e: // EOR
				u.handleEor(mode)
			case 0x1f: // SFT
				u.handleSft(mode)
			}
		}
	}
	return false
}