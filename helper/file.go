package helper

import (
	"os"
	"strconv"
	"strings"
)

func GetFileExtension(filename string) string {
	arr := strings.Split(filename, ".")
	return arr[len(arr)-1]
}

func AllowedExtension() map[string]bool {
	allow := make(map[string]bool)
	allow["png"] = true
	allow["jpg"] = true
	allow["jpeg"] = true
	return allow
}

func WriteFile(path string, fileByte []byte) {
	f, err := os.Create(path)
	PanicIfError(err)
	defer f.Close()

	_, err = f.Write(fileByte)
	PanicIfError(err)

	err = f.Sync()
	PanicIfError(err)
}

func ReadFile(path string) []byte {
	fileBytes, err := os.ReadFile(path)
	PanicIfError(err)
	return fileBytes
}

func GenerateImageFilename() string {
	randomStr := GenerateRandomString(5)
	unixNano := strconv.Itoa(GetUnixNano())
	ext := ".png"
	return randomStr + "_" + unixNano + ext
}
