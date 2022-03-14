package util

import "strconv"

//int64 to string
func TranToString(i int64) string {
	return strconv.FormatInt(i, 10)
}
