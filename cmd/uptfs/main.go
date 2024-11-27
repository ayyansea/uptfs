package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"slices"

	"github.com/alexflint/go-arg"
	"github.com/ayyansea/uptfs/internal/config"
	"github.com/ayyansea/uptfs/internal/filter"
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

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	inputString := scanner.Text()

	if inputString == "" {
		err := errors.New("the input string is empty")
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}

	additionalDelimeters := []string{",", ".", " "}
	tempword := ""
	var tokenlist token.LinkedTokenList

	for index, character := range inputString {
		if slices.Contains(additionalDelimeters, string(character)) {
			if len(tempword) != 0 {
				for _, filterName := range config.Filters {
					currentfilter := filter.FilterList[filterName]()
					tempword = currentfilter.Filter(tempword)
				}
				tokenlist.AddToken(tempword)
				tokenlist.AddToken(string(character))
				tempword = ""

				continue
			}
			tokenlist.AddToken(string(character))
			tempword = ""

			continue
		}

		tempword = tempword + string(character)
		if index == len(inputString)-1 {
			for _, filterName := range config.Filters {
				currentfilter := filter.FilterList[filterName]()
				tempword = currentfilter.Filter(tempword)
			}
			tokenlist.AddToken(tempword)
			tempword = ""
		}
	}

	result := ""

	for current := tokenlist.GetHead(); current != nil; current = current.GetNextToken() {
		result = result + current.GetContent()
	}

	fmt.Println(result)
}
