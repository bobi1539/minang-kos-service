package helper

import (
	"strconv"
	"strings"
)

func StringQueryLike(value string) string {
	return "%" + strings.ToLower(value) + "%"
}

func StringToInt(value string) int {
	intValue, err := strconv.Atoi(value)
	PanicIfError(err)
	return intValue
}

func GetSqlOffset(page int, size int) int {
	return (page - 1) * size
}
