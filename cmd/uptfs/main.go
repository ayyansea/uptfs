package main

import (
	"bufio"
	"errors"
	"fmt"
	"log/slog"
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
	Verbose    bool   `arg:"-v" help:"toggle verbose (debug) mode"`
}

func errExit(err error) {
	fmt.Println(err)
	os.Exit(1)
}

func main() {
	arg.MustParse(&args)

	if args.Verbose {
		slog.SetLogLoggerLevel(slog.LevelDebug)
	}

	slog.Debug("uptfs started")

	var configFilePath string
	var err error

	if args.ConfigFile != "" {
		configFilePath, err = filepath.Abs(args.ConfigFile)
		slog.Debug("config file path: " + configFilePath)
	}
	if err != nil {
		errExit(err)
	}

	var config config.Config
	config.LoadConfig(configFilePath)

	var inputString string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		inputString = inputString + scanner.Text()
	}

	if inputString == "" {
		err := errors.New("the input string is empty")
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}

	additionalDelimeters := []string{",", ".", " "}
	var delimeterString string
	for _, element := range additionalDelimeters {
		if element == " " {
			element = "<space>"
		}
		delimeterString = delimeterString + element + " "
	}
	slog.Debug("delimeters: " + delimeterString)

	tempword := ""
	var tokenlist token.LinkedTokenList

	for index, character := range inputString {
		slog.Debug("current character: " + string(character))
		if slices.Contains(additionalDelimeters, string(character)) {
			slog.Debug("current character is a delimeter")
			if len(tempword) != 0 {
				slog.Debug("word is longer than 0 characters, filtering it")
				for _, filterName := range config.Filters {
					_, ok := filter.FilterList[filterName]
					if ok {
						slog.Debug("using filter " + filterName)
						currentfilter := filter.FilterList[filterName]()
						tempword = currentfilter.Filter(tempword)
					} else {
						slog.Debug("there's no filter named " + filterName)
						continue
					}
				}
				slog.Debug("filtered word: " + tempword)
				tokenlist.AddToken(tempword)
				slog.Debug("adding delimeter")
				tokenlist.AddToken(string(character))
				slog.Debug("resetting current word to an empty string")
				tempword = ""

				continue
			}
			slog.Debug("current word is a zero-length string")
			slog.Debug("adding delimeter")
			tokenlist.AddToken(string(character))
			tempword = ""

			continue
		}
		slog.Debug("current character is not a delimeter, adding to word")
		tempword = tempword + string(character)
		if index == len(inputString)-1 {
			slog.Debug("string ended, filtering current word")
			for _, filterName := range config.Filters {
				_, ok := filter.FilterList[filterName]
				if ok {
					slog.Debug("using filter " + filterName)
					currentfilter := filter.FilterList[filterName]()
					tempword = currentfilter.Filter(tempword)
				} else {
					slog.Debug("there's no filter named " + filterName)
					continue
				}
			}
			slog.Debug("filtered word: " + tempword)
			tokenlist.AddToken(tempword)
			tempword = ""
		}
	}

	result := ""

	for current := tokenlist.GetHead(); current != nil; current = current.GetNextToken() {
		result = result + current.GetContent()
	}

	fmt.Println(result)

	slog.Debug("uptfs finished")
}
