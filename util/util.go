package util

import (
	"fmt"
	"math"
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

func GetPtr[T string | int](value T) *T {
	return &value
}

func CloneMap[T comparable, S any](dst map[T]S, src map[T]S) {
	for t, s := range src {
		dst[t] = s
	}
}

func Reverse[T any](numbers []T) []T {
	newNumbers := make([]T, 0, len(numbers))
	for i := len(numbers) - 1; i >= 0; i-- {
		newNumbers = append(newNumbers, numbers[i])
	}
	return newNumbers
}

func FloydWarshall(matrix [][]int) [][]int {
	n := len(matrix)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = matrix[i][j]
		}
	}

	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if dp[i][k] == math.MaxInt || dp[k][j] == math.MaxInt {
					continue
				}
				if dp[i][j] > dp[i][k]+dp[k][j] {
					dp[i][j] = dp[i][k] + dp[k][j]
				}
			}
		}
	}
	return dp
}
