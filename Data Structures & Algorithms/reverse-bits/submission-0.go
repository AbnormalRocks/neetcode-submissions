func reverseBits(n int) int {
	x := uint32(n) // treat as 32-bit unsigned pattern
	var result uint32
	for i := 0; i < 32; i++ {
		result <<= 1
		result |= x & 1
		x >>= 1
	}
	return int(result)
}