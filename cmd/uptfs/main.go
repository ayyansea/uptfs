package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/alexflint/go-arg"
	"github.com/ayyansea/uptfs/internal/config"
	"github.com/ayyansea/uptfs/internal/filter"
	"github.com/ayyansea/uptfs/internal/split"
	"github.com/ayyansea/uptfs/internal/token"
)

var args struct {
	ConfigFile string `arg:"-c" help:"path to config file" default:""`
}

func errExit(err error) {
	fmt.Println(err)
	os.Exit(1)
}

func main() {
	arg.MustParse(&args)

	var configFilePath string
	var err error

	if args.ConfigFile != "" {
		configFilePath, err = filepath.Abs(args.ConfigFile)
	}
	if err != nil {
		errExit(err)
	}

	var config config.Config
	config.LoadConfig(configFilePath)
	fmt.Printf("Config: %v\n", config)

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

	for current := linkedTokens.GetHead(); current != nil; current = current.GetNextToken() {
		for _, filterName := range config.Filters {
			filter := filter.FilterList[filterName]()
			current.SetContent(filter.Filter(current.GetContent()))
		}
		fmt.Println(current.GetContent())
	}
}
