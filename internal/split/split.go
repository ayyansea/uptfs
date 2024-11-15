package split

import (
	"strings"
)

func checkArrayElementsEmpty(array []string) bool {
	for _, str := range array {
		if str != "" {
			return false
		}
	}
	return true
}

func NewArrayWithSplit(initialArray []string, index int, token string, delimeter string) (result []string) {
	split := strings.Split(token, delimeter)
	splitLength := len(split)

	/*
		When a token only consists of delimeter * N (N >= 0),
		the resulting split consists of N empty elements.
		Here we check if it is so and essentialy remove that token
		from resulting array.
	*/
	splitIsEmpty := checkArrayElementsEmpty(split)
	if splitIsEmpty {
		result = append(initialArray[:index], initialArray[index+1:]...)
		return result
	}

	if splitLength > 1 {
		if split[splitLength-1] == "" {
			split = split[:splitLength-1]
		}

		result = append(initialArray[:index], append(split, initialArray[index+1:]...)...)
	}
	if splitLength == 1 {
		result = initialArray
	}

	return result
}
