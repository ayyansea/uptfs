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

	if len(tokens) == 0 {
		err := errors.New("the slice is empty")
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}

	var linkedTokens LinkedTokenList
	SliceToLinkedTokenSlice(tokens, &linkedTokens)

	lowercaseFilter := NewLowercaseFilterWithExtraSteps()
	uppercaseFilter := NewUppercaseFilter()

	for current := linkedTokens.head; current != nil; current = current.GetNextToken() {
		current.content = lowercaseFilter.Filter(current.content)
		current.content = uppercaseFilter.Filter(current.content)
	}

	fmt.Println("Head: ", linkedTokens.head.content)
	fmt.Println("Tail: ", linkedTokens.tail.content)
}

func formatInput(tokenArray []string, delimeterArray []string) []string {
	for _, delimeter := range delimeterArray {
		for index, element := range tokenArray {
			tokenArray = split.NewArrayWithSplit(tokenArray, index, element, delimeter)
		}
	}

	return tokenArray
}
