package luncheon

type Util struct{}

// Used for easy conversion of uint32 and []byte
func (u Util) Uint32toB(val uint32) []byte { // TODO: make not copywritted
	output := make([]byte, 4)
	for i := uint32(0); i < 4; i++ {
		output[i] = byte((val >> (8 * i)) & 0xff)
	}
	return output
}
