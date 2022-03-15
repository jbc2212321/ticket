package util

import (
	"net/http"
	"strconv"
)

//int64 to string
func TranToString(i int64) string {
	return strconv.FormatInt(i, 10)
}

//string to int64
func TranToInt64(s string) int64 {
	i, _ := strconv.ParseInt(s, 10, 64)
	return i
}

func GetResponse() Response {
	return Response{
		Status: http.StatusOK,
		Code:   0,
	}
}
