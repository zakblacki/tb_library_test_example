package lib

//export add
func add(a, b uint32) uint64 {
	return uint64(a) + uint64(b)
}
