package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/alexflint/go-arg"
	"github.com/ayyansea/uptfs/internal/filter"
	"github.com/ayyansea/uptfs/internal/split"
	"github.com/ayyansea/uptfs/internal/token"
)

var args struct {
	Foo string `help:"it's a foo"`
	Bar bool   `help:"it's a bar"`
}

func main() {
	arg.MustParse(&args)
	fmt.Println(args.Foo, args.Bar)

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
	tokens = split.FormatInput(tokens, additionalDelimeters)

	if len(tokens) == 0 {
		err := errors.New("the slice is empty")
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}

	var linkedTokens token.LinkedTokenList
	token.SliceToLinkedTokenSlice(tokens, &linkedTokens)

	lowercaseFilter := filter.NewLowercaseFilterWithExtraSteps()
	uppercaseFilter := filter.NewUppercaseFilter()

	for current := linkedTokens.GetHead(); current != nil; current = current.GetNextToken() {
		current.SetContent(lowercaseFilter.Filter(current.GetContent()))
		current.SetContent(uppercaseFilter.Filter(current.GetContent()))
	}

	head := linkedTokens.GetHead()
	tail := linkedTokens.GetTail()
	fmt.Println("Head: ", head.GetContent())
	fmt.Println("Tail: ", tail.GetContent())
}
