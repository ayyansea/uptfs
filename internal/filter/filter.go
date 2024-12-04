package filter

type Filter struct {
	name       string
	action     func(string) string
	subfilters []Filter
}

type Filterer interface {
	Filter(string) string
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
