package p9

func Run(x int) bool {
	return isPalindrome(x)
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	if x != 0 && x%10 == 0 {
		return false
	}

	revX := 0

	for x > revX {
		revX = revX*10 + x%10
		x = x / 10
	}

	if x == revX || x == revX/10 {
		return true
	}

	return false
}
