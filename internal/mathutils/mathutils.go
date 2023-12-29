package mathutils

func Sign(i int) int {
	if i == 0 {
		return 0
	}
	if i > 0 {
		return 1
	}
	return -1
}
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
