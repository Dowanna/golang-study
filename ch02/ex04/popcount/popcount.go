package popcount

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
	// fmt.Fprintf(os.Stdout, "%.08b\n", byte(128>>2))
}

func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func PopCountNew(x uint64) int {
	counter := 0

	for i := 1; i < 65; i++ {
		if int(byte(x&1)) == 1 {
			counter++
		}
		x = x >> uint(1)
	}

	return counter
}
