package cmd

type LinkedTokenList struct {
	head, tail *Token
}

func (lts *LinkedTokenList) AddToken(content string) {
	newToken := &Token{
		content: content,
		prev:    nil,
		next:    nil,
	}

	if lts.head == nil {
		lts.head = newToken
		lts.tail = newToken
	} else {
		newToken.SetPreviousToken(lts.tail)
		lts.tail.SetNextToken(newToken)
		lts.tail = newToken
	}
}

func SliceToLinkedTokenSlice(slice []string, tokenSlice *LinkedTokenList) {
	for _, item := range slice {
		tokenSlice.AddToken(item)
	}
}
