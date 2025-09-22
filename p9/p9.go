package p9

func Run(x int) bool {
	return isPalindrome(x)
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	xTab := []int{}

	for x != 0 {
		tmp := x % 10
		xTab = append(xTab, tmp)
		x = x / 10
	}

	xLen := len(xTab)

	if xLen <= 1 {
		return true
	}

	for key, val := range xTab {
		tmp := xLen - (key + 1)
		if xTab[tmp] != val {
			return false
		}
	}

	return true
}
