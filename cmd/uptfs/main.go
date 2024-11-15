package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
	"uptfs/internal/config"
	"uptfs/internal/split"
)

const filepath = ""

func main() {
	var config config.Config
	config.LoadConfig(filepath)

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

	fmt.Println(tokens)
}

func formatInput(tokenArray []string, delimeterArray []string) []string {
	for _, delimeter := range delimeterArray {
		for index, element := range tokenArray {
			tokenArray = split.NewArrayWithSplit(tokenArray, index, element, delimeter)
		}
	}

	return tokenArray
}
