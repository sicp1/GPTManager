package Tool

import "strconv"

func StrToInt(str string) int {
	temp, _ := strconv.Atoi(str)
	return temp
}
