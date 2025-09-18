package p7

func Run(x int) int {
	return reverse(x)
}

func reverse(x int) int {
	reversed := 0

	if x < -2147483648 || x > 2147483647 {
		return 0
	}

	for x != 0 {
		remained := x % 10
		reversed = reversed*10 + remained
		x = x / 10
	}

	if reversed < -2147483648 || reversed > 2147483647 {
		return 0
	}
	return reversed
}
