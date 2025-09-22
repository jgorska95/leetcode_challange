package p10

func Run(s string, p string) bool {
	return isMatch(s, p)
}

func isMatch(s string, p string) bool {
	if p == ".*" {
		return true
	}
	ptrP := 0

	for _, val := range s {
		if ptrP >= len(p) {
			return false
		}

		if p[ptrP] == '*' {
			if ptrP == 0 {
				ptrP++
				continue
			}
			if p[ptrP-1] == "." || (p[ptrP-1] == val) {
				continue
			}
		}
		if val == p[ptrP] || p[ptrP] == "." {
			ptrP++
			continue
		}
		return false
	}

	return true
}

func helper(s, p string, i, j int) {

}
