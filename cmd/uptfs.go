package cmd

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
	"uptfs/util/split"
)

func Run() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	inputString := scanner.Text()

	if inputString == "" {
		err := errors.New("the input string is empty")
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}

	additionalDelimeters := []string{",", "."}
	tokens := strings.Split(inputString, " ")
	tokens = formatInput(tokens, additionalDelimeters)

	lowercaseFilter := NewLowercaseFilterWithExtraSteps()
	lowercaseFilter.Greet()

	var newTokens []string
	for _, token := range tokens {
		newTokens = append(newTokens, lowercaseFilter.Filter(token))
	}

	tokens = newTokens

	uppercaseFilter := NewUppercaseFilter()
	uppercaseFilter.Greet()

	newTokens = nil
	for _, token := range tokens {
		newTokens = append(newTokens, uppercaseFilter.Filter(token))
	}

	result := strings.Join(newTokens, " ")

	fmt.Println("")
	fmt.Println(result)
}

func formatInput(tokenArray []string, delimeterArray []string) []string {
	for _, delimeter := range delimeterArray {
		for index, element := range tokenArray {
			tokenArray = split.NewArrayWithSplit(tokenArray, index, element, delimeter)
		}
	}

	return tokenArray
}
