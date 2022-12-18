package util

import (
	"fmt"
	"os"
	"strings"
)

func PanicOnError(err error) {
	if err != nil {
		panic(err)
	}
}

func GetFileContentsSplit(filename string) ([]string, error) {
	content, err := os.ReadFile(filename) // the file is inside the local directory
	if err != nil {
		fmt.Println("Err", err)
		return nil, err
	}
	fullFile := string(content)
	fileSplit := strings.Split(fullFile, "\n")
	return fileSplit, nil
}

func Contains[T comparable](s []T, e T) bool {
	for _, v := range s {
		if v == e {
			return true
		}
	}
	return false
}
