package textnode

import (
	"ntoolkit/parser"
	"strings"
)

type wordTokenizer struct {
	next   *parser.Token
	tokens *parser.Tokens
}

// Enter starts processing a data stream with a context object and resets the internal tokenizer state
func (t *wordTokenizer) Enter(context *parser.Tokens) {
	t.tokens = context
	t.next = nil
}

// Process reads incoming string data and pushes tokens to the context
func (t *wordTokenizer) Process(data string) {
	parts := strings.Split(data, "")
	for i := range parts {
		if token := t.process(parts[i]); token != nil {
			t.tokens.Push(token)
		}
	}
}

// Close should close and resolve the tokenizer.
func (t *wordTokenizer) Close() {
	if t.next != nil {
		t.tokens.Push(t.next)
		t.next = nil
	}
}

// Process reads incoming string data and pushes tokens to the context
func (t *wordTokenizer) process(sym string) *parser.Token {
	// Spaces always end the current token, if there is one, and leave no pending token
	if sym == " " {
		if t.next != nil {
			rtn := t.next
			t.next = nil
			return rtn
		}
		return nil
	}

	// If there's nothing on the stack, create a new symbol
	if t.next == nil {
		t.next = &parser.Token{Type: parser.TokenTypeNone, Raw: &sym}
		return nil
	}

	// If the top of the stack already exists and is None type, append to it
	if t.next != nil && t.next.Type == parser.TokenTypeNone {
		*t.next.Raw += sym
		return nil
	}

	// No idea how we might get here, do nothing.
	return nil
}
