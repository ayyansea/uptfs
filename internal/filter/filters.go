package filter

import (
	"strings"
)

var FilterList = map[string]func() Filter{
	"uppercase":       NewUppercaseFilter,
	"lowercase":       NewLowercaseFilter,
	"lowercase_extra": NewLowercaseFilterWithExtraSteps,
	"doubler":         NewDoubleFilter,
	"inverse":         NewInverseFilter,
	"l33t":            NewLeetFilter,
	"reverse":         NewReverseFilter,
	"uwu":             NewUwuFilter,
}

func NewUppercaseFilter() Filter {
	uppercaseFilter := Filter{
		name:       "Uppercase",
		action:     strings.ToUpper,
		subfilters: []Filter{},
	}

	return uppercaseFilter
}

func NewLowercaseFilter() Filter {
	lowercaseFilter := Filter{
		name:       "Lowercase",
		action:     strings.ToLower,
		subfilters: []Filter{},
	}

	return lowercaseFilter
}

func NewLowercaseFilterWithExtraSteps() Filter {
	var subfilters []Filter

	subfilters = append(subfilters, NewUppercaseFilter())
	subfilters = append(subfilters, NewLowercaseFilter())

	filter := Filter{
		name:       "Lowercase (extra dumb)",
		action:     func(string) string { return "" },
		subfilters: subfilters,
	}

	return filter
}

func doubleFilterAction(content string) string {
	var temp strings.Builder

	for _, character := range content {
		for i := 0; i < 2; i++ {
			temp.WriteString(string(character))
		}
	}

	return temp.String()
}

func NewDoubleFilter() Filter {
	filter := Filter{
		name:       "Doubler",
		action:     doubleFilterAction,
		subfilters: []Filter{},
	}

	return filter
}

func inverseFilterAction(content string) string {
	var temp strings.Builder

	for _, character := range content {
		strchar := string(character)
		lower := strings.ToLower(strchar)
		upper := strings.ToUpper(strchar)

		if strchar == lower {
			temp.WriteString(upper)
		}

		if strchar == upper {
			temp.WriteString(lower)
		}
	}

	return temp.String()
}

func NewInverseFilter() Filter {
	filter := Filter{
		name:       "Inverse case",
		action:     inverseFilterAction,
		subfilters: []Filter{},
	}

	return filter
}

func leetFilterAction(content string) string {
	replacements := map[string]string{
		"e": "3",
		"E": "3",
		"i": "1",
		"I": "1",
		"a": "4",
		"A": "4",
		"b": "8",
		"B": "8",
		"t": "7",
		"T": "7",
		"o": "0",
		"O": "0",
	}

	var temp strings.Builder

	for _, character := range content {
		strchar := string(character)
		value, ok := replacements[strchar]
		if ok {
			temp.WriteString(value)
		} else {
			temp.WriteString(strchar)
		}
	}

	return temp.String()
}

func NewLeetFilter() Filter {
	filter := Filter{
		name:       "l33t",
		action:     leetFilterAction,
		subfilters: []Filter{},
	}

	return filter
}

// https://stackoverflow.com/a/1754209
func reverseString(content string) string {
	n := 0
	rune := make([]rune, len(content))
	for _, r := range content {
		rune[n] = r
		n++
	}
	rune = rune[0:n]

	for i := 0; i < n/2; i++ {
		rune[i], rune[n-1-i] = rune[n-1-i], rune[i]
	}

	output := string(rune)
	return output
}

func NewReverseFilter() Filter {
	filter := Filter{
		name:       "Reverser",
		action:     reverseString,
		subfilters: []Filter{},
	}

	return filter
}

func uwuFilterAction(content string) string {
	var temp strings.Builder

	replacements := map[string]string{
		"r": "w",
		"u": "uwu",
		"l": "w",
		"R": "W",
		"U": "UWU",
		"L": "W",
	}

	for _, character := range content {
		strchar := string(character)
		value, ok := replacements[strchar]
		if ok {
			temp.WriteString(value)
		} else {
			temp.WriteString(strchar)
		}
	}

	return temp.String()
}

func NewUwuFilter() Filter {
	filter := Filter{
		name:       "Uwuifier",
		action:     uwuFilterAction,
		subfilters: []Filter{},
	}

	return filter
}
