package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/alexflint/go-arg"
	"github.com/ayyansea/uptfs/internal/filter"
	"github.com/ayyansea/uptfs/internal/split"
	"github.com/ayyansea/uptfs/internal/token"
)

var args struct {
	ConfigFile string   `arg:"-c" help:"path to config file" default:""`
	InputFile  string   `arg:"-i" help:"path to input file" default:""`
	OutputFile string   `arg:"-o" help:"path to output file" default:""`
	Filters    []string `arg:"-f" help:"list of filters" default:""`
}

func errExit(err error) {
	fmt.Println(err)
	os.Exit(1)
}

func main() {
	arg.MustParse(&args)

	var configFilePath, inputFilePath, outFilePath string
	var err error

	if args.ConfigFile != "" {
		configFilePath, err = filepath.Abs(args.ConfigFile)
	}
	if err != nil {
		errExit(err)
	}

	if args.InputFile != "" {
		inputFilePath, err = filepath.Abs(args.InputFile)
	}
	if err != nil {
		errExit(err)
	}

	if args.OutputFile != "" {
		outFilePath, err = filepath.Abs(args.OutputFile)
	}
	if err != nil {
		errExit(err)
	}

	fmt.Printf("%v %v %v\n",
		configFilePath,
		inputFilePath,
		outFilePath)

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
}
