package emulator

type Ram struct {
	workRam [8192]byte
}

func (ram *Ram) load(address uint16) byte {
	switch {
	case (address >= 0) && (address <= 3):
		return 1
	default:
		return 2
	}
}
