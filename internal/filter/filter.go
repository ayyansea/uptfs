package filter

import (
	"fmt"
)

type Filter struct {
	name       string
	action     func(string) string
	subfilters []Filter
}

type Greeter interface {
	Greet()
}

type Filterer interface {
	Greeter
	Filter(string) string
}

func (f Filter) Greet() {
	subfilterCount := len(f.subfilters)

	fmt.Printf("I am a filter and my name is %v\n", f.name)
	if subfilterCount > 0 {
		fmt.Println("My subfilters are:")

		for _, subfilter := range f.subfilters {
			fmt.Printf("- %v\n", subfilter.name)
		}
	}
}

func (f Filter) Filter(token string) (modifiedToken string) {
	subfilterCount := len(f.subfilters)
	modifiedToken = token

	if subfilterCount > 0 {
		for _, subfilter := range f.subfilters {
			modifiedToken = subfilter.action(modifiedToken)
		}

		return modifiedToken
	}

	modifiedToken = f.action(token)
	return modifiedToken
}
