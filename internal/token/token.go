package token

type Token struct {
	content    string
	prev, next *Token
}

func (t *Token) GetContent() string {
	return t.content
}

func (t *Token) SetContent(content string) {
	t.content = content
}

func (t *Token) GetPreviousToken() *Token {
	return t.prev
}

func (t *Token) SetPreviousToken(newToken *Token) {
	t.prev = newToken
}

func (t *Token) GetNextToken() *Token {
	return t.next
}

func (t *Token) SetNextToken(newToken *Token) {
	t.next = newToken
}
