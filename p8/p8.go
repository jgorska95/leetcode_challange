package p8

import (
	"strconv"
	"strings"
)

func Run(s string) int {
	return myAtoi(s)
}

func myAtoi(s string) int {
	retVal := 0
	sign := 1
	ignore := true
	started := false

	noSpace := strings.Fields(s)
	if len(noSpace) < 1 {
		return retVal
	}

	stringSlice := strings.Split(noSpace[0], "")

	for key, val := range stringSlice {
		if stringSlice[key] == "-" {
			if started {
				break
			}
			sign = -1
			started = true
			continue
		}
		if stringSlice[key] == "+" {
			if started {
				break
			}
			started = true
			continue
		}

		tmp, err := strconv.Atoi(val)
		if err != nil {
			break
		}
		if tmp == 0 && ignore {
			started = true
			continue
		}
		retVal = retVal*10 + tmp
		ignore = false
		started = true

		if retVal < -2147483648 || retVal > 2147483647 {
			limit := 2147483647
			if sign < 0 {
				limit++
			}
			return limit * sign
		}
	}

	retVal = retVal * sign

	return retVal
}
