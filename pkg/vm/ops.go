package vm

// Helper functions for opcode handlers
func is2(mode uint8) bool {
	return mode&0x20 != 0
}

func isReturn(mode uint8) bool {
	return mode&0x40 != 0
}

func isKeep(mode uint8) bool {
	return mode&0x80 != 0
}

// Opcode handlers

func (u *Uxn) handleInc(mode uint8) {
	if isReturn(mode) {
		a := u.rst.dec()
		if is2(mode) {
			a |= uint8(u.rst.dec()) << 8
			u.rst.dat[u.rst.ptr] = uint8((uint16(a) + 1) >> 8)
			u.rst.ptr++
		}
		u.rst.dat[u.rst.ptr] = uint8(a + 1)
		u.rst.ptr++
	} else {
		a := u.wst.dec()
		if is2(mode) {
			a |= uint8(u.wst.dec()) << 8
			u.wst.dat[u.wst.ptr] = uint8((uint16(a) + 1) >> 8)
			u.wst.ptr++
		}
		u.wst.dat[u.wst.ptr] = uint8(a + 1)
		u.wst.ptr++
	}
}

func (u *Uxn) handlePop(mode uint8) {
	if isReturn(mode) {
		if is2(mode) {
			u.rst.ptr -= 2
		} else {
			u.rst.ptr--
		}
	} else {
		if is2(mode) {
			u.wst.ptr -= 2
		} else {
			u.wst.ptr--
		}
	}
}

func (u *Uxn) handleSwp(mode uint8) {
	if isReturn(mode) {
		x := u.rst.dec()
		y := u.rst.dec()
		if is2(mode) {
			x2 := u.rst.dec()
			y2 := u.rst.dec()
			u.rst.dat[u.rst.ptr] = x2
			u.rst.ptr++
			u.rst.dat[u.rst.ptr] = x
			u.rst.ptr++
			u.rst.dat[u.rst.ptr] = y2
			u.rst.ptr++
			u.rst.dat[u.rst.ptr] = y
			u.rst.ptr++
		} else {
			u.rst.dat[u.rst.ptr] = x
			u.rst.ptr++
			u.rst.dat[u.rst.ptr] = y
			u.rst.ptr++
		}
	} else {
		x := u.wst.dec()
		y := u.wst.dec()
		if is2(mode) {
			x2 := u.wst.dec()
			y2 := u.wst.dec()
			u.wst.dat[u.wst.ptr] = x2
			u.wst.ptr++
			u.wst.dat[u.wst.ptr] = x
			u.wst.ptr++
			u.wst.dat[u.wst.ptr] = y2
			u.wst.ptr++
			u.wst.dat[u.wst.ptr] = y
			u.wst.ptr++
		} else {
			u.wst.dat[u.wst.ptr] = x
			u.wst.ptr++
			u.wst.dat[u.wst.ptr] = y
			u.wst.ptr++
		}
	}
}

// Additional opcode handlers would be implemented here...