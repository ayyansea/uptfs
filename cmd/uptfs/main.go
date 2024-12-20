package main

import (
	"bufio"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/alexflint/go-arg"
	"github.com/ayyansea/uptfs/internal/config"
	"github.com/ayyansea/uptfs/internal/filter"
	"github.com/ayyansea/uptfs/internal/token"
)

var args struct {
	ConfigFile string   `arg:"-c" help:"path to config file" default:""`
	Verbose    bool     `arg:"-v" help:"toggle verbose (debug) mode"`
	InputFile  string   `arg:"-i" help:"name of input file to read text from"`
	OutputFile string   `arg:"-o" help:"name of output file to write text to"`
	Filter     []string `arg:"-f,separate" help:"name of a filter that will be applied to text, can be specified multiple times"`
}

func main() {
	arg.MustParse(&args)

	if args.Verbose {
		slog.SetLogLoggerLevel(slog.LevelDebug)
	}

	slog.Debug("uptfs started")

	var configFilePath string

	var config config.Config
	var err error

	if args.ConfigFile != "" {
		configFilePath, err = filepath.Abs(args.ConfigFile)
		config.LoadConfig(configFilePath)
		slog.Debug("config file path: " + configFilePath)
	}
	if err != nil {
		slog.Error(err.Error())
	}

	config.Filters = append(config.Filters, args.Filter...)

	var scanner *bufio.Scanner
	var inputFileData []byte
	var inputStrings []string

	if len(args.InputFile) > 0 {
		inputFilePath, _ := filepath.Abs(args.InputFile)
		inputFileData, err = os.ReadFile(inputFilePath)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		scanner = bufio.NewScanner(strings.NewReader(string(inputFileData)))
		for scanner.Scan() {
			inputStrings = append(inputStrings, scanner.Text())
		}
	} else {
		scanner = bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			inputStrings = append(inputStrings, scanner.Text())
		}
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
	var tokenizedInputSlice []token.LinkedTokenList
	var tokenlist token.LinkedTokenList

	for _, elem := range inputStrings {
		for index, character := range elem {
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
			if index == len(elem)-1 {
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
		tokenizedInputSlice = append(tokenizedInputSlice, tokenlist)
		tokenlist.Clear()
	}

	if len(args.OutputFile) > 0 {
		var result []byte

		outputFilePath, _ := filepath.Abs(args.OutputFile)
		os.Create(outputFilePath)

		for _, elem := range tokenizedInputSlice {
			for current := elem.GetHead(); current != nil; current = current.GetNextToken() {
				result = append(result, []byte(current.GetContent())...)
			}
			result = append(result, []byte("\n")...)
			f, _ := os.OpenFile(outputFilePath, os.O_APPEND|os.O_WRONLY, 0644)
			defer f.Close()
			if _, err = f.Write(result); err != nil {
				panic(err)
			}
			result = []byte("")
		}
	} else {
		var result string

		for _, elem := range tokenizedInputSlice {
			for current := elem.GetHead(); current != nil; current = current.GetNextToken() {
				result = result + current.GetContent()
			}
			fmt.Println(result)
			result = ""
		}
	}

	slog.Debug("uptfs finished")
}
