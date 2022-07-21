package basic

var sum int

func addition(args ...int) int {
	for _, v := range args {
		sum += v
	}
	return sum
}
