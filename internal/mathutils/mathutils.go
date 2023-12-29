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

func MeanF(data []float64) float64 {
	if len(data) == 0 {
		return 0
	}
	var sum float64
	for _, d := range data {
		sum += d
	}
	return sum / float64(len(data))
}

func MeanI(data ...int) int {
	if len(data) == 0 {
		return 0
	}
	var sum int
	for _, d := range data {
		sum += d
	}
	return sum / len(data)
}
