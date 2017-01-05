package textnode

import (
	"ntoolkit/iter"
	"ntoolkit/errors"
	"fmt"
	"unicode/utf8"
)

// Text is the lowest level representation of a string that the core can return.
type Text struct {
	Value  string
	Styles []*Style
	Errors []error
}

// TextToken is a single continuous stream of runes with the same style applied to them.
type TextToken struct {
	Value string
	Style *Style
}

// newText returns a new Text object for the given values
func newText(text string, style string, styleMap map[rune]string, env *Env) *Text {
	styleCount := utf8.RuneCountInString(style)
	rtn := Text{Value: text, Styles: make([]*Style, styleCount), Errors: make([]error, 0)}
	styleRunes := []rune(style)
	for i := 0; i < styleCount; i++ {
		rtn.Styles[i] = &env.Stylesheet.Default
		if styleId, ok := styleMap[styleRunes[i]]; ok {
			rtn.Styles[i] = env.Stylesheet.Get(styleId)
			if rtn.Styles[i] == nil {
				rtn.Styles[i] = &env.Stylesheet.Default
				rtn.Errors = append(rtn.Errors, errors.Fail(ErrBadStyles{}, nil, fmt.Sprintf("No entry in style sheet for style id '%s'", styleId)))
			}
		} else {
			// Always allow ' ' as empty padding
			if styleRunes[i] != ' ' {
				rtn.Errors = append(rtn.Errors, errors.Fail(ErrBadStyles{}, nil, fmt.Sprintf("No entry in style map for rune '%c'", styleRunes[i])))
			}
		}
	}
	return &rtn
}

// Tokens returns a text object as an iterator of TextTokens
func (t *Text) Tokens() iter.Iter {
	return newTextTokenIter(t)
}

// Render a text stream using a TokenRenderer and return the combined result
func (t *Text) Render(renderer TokenRenderer) (string, error) {

	// First render each token as a styled block and combine the result
	tokens := newTextTokenIter(t)
	var iterErr error = nil
	var val interface{} = nil
	var combined = ""
	for val, iterErr = tokens.Next(); iterErr == nil; val, iterErr = tokens.Next() {
		styled, err := renderer.AsString(val.(TextToken))
		if err != nil {
			return "", err
		}
		combined += styled
	}
	if !errors.Is(iterErr, iter.ErrEndIteration{}) {
		return "", iterErr
	}

	return combined, nil
}