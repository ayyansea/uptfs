package token

type LinkedTokenList struct {
	head, tail *Token
}

func (lts *LinkedTokenList) GetHead() *Token {
	return lts.head
}

func (lts *LinkedTokenList) GetTail() *Token {
	return lts.tail
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
