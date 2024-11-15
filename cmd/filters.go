package cmd

import (
	"strings"
)

var filterList = map[string]interface{}{
	"uppercase":       NewUppercaseFilter,
	"lowercase":       NewLowercaseFilter,
	"lowercase_extra": NewLowercaseFilterWithExtraSteps,
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
