package luncheon

// Used for easy conversion of uint32 and []byte
func Uint32toB(val uint32) []byte { // TODO: make not copywritted
	r := make([]byte, 4)
	for i := uint32(0); i < 4; i++ {
		r[i] = byte((val >> (8 * i)) & 0xff)
	}
	return r
}
