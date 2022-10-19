package tools

import "strconv"

func StrToInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func StrToInt64(s string) int64 {
	i, _ := strconv.ParseInt(s, 10, 64)
	return i
}
