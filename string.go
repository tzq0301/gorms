package gorms

import (
	"golang.org/x/exp/constraints"
	"strconv"
)

// SignedIntegerToString convert signed integer (int, int8, int16, int32, int64) to string
//
// Refer to https://stackoverflow.com/questions/39442167/convert-int32-to-string-in-golang
func SignedIntegerToString[T constraints.Signed](i T) string {
	return strconv.FormatInt(int64(i), 10)
}
