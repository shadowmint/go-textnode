package textnode

import (
	"ntoolkit/parser"
)

type Renderer interface {
	// Render tokens into device string
	Render(tokens *parser.Tokens) (string, error)
}

type PassThroughRenderer struct {
}

func (r *PassThroughRenderer) Render(tokens *parser.Tokens) (string, error) {
	if tokens == nil {
		return "", nil
	}
	if tokens.Count() == 0 {
		return "", nil
	}
	return tokens.Front.WalkRaw(" "), nil
}
