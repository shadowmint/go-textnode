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

// TextTokenIter is an iterator that returns blocks of text
type TextTokenIter struct {
	error  error
	offset int
	length int
	text   *Text
	runes  []rune
}

// TextRenderer renders a text object out as a string
type TextRenderer interface {
	// AsString should return the appropriate device formatted string for the given text value.
	AsString(text *Text) string
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
			if styleRunes[i] != ' ' {  // Always allow ' ' as empty padding
				rtn.Errors = append(rtn.Errors, errors.Fail(ErrBadStyles{}, nil, fmt.Sprintf("No entry in style map for rune '%c'", styleRunes[i])))
			}
		}
	}
	return &rtn
}

// Tokens returns a text object as an iterator of TextTokens
func (t *Text) Tokens() iter.Iter {
	runes := []rune(t.Value)
	return &TextTokenIter{
		error: nil,
		runes: runes,
		length: len(runes),
		offset: 0,
		text:t}
}

func (iterator *TextTokenIter) Next() (interface{}, error) {
	if iterator.error != nil {
		return nil, iterator.error
	}

	// Bounds check
	if len(iterator.text.Styles) != len(iterator.runes) {
		iterator.error = errors.Fail(ErrBadStyles{}, nil, fmt.Sprintf("Incorrect number of styles (%d) for string of length %d", len(iterator.text.Styles), len(iterator.runes)))
	}

	// Find next sequence
	value := ""
	style := iterator.text.Styles[iterator.offset]
	for i := iterator.offset; i < len(iterator.runes); i++ {
		if iterator.text.Styles[i] != style {
			break
		} else {
			value += string(iterator.runes[i])
			iterator.offset = i + 1
		}
	}

	// End sequence?
	if iterator.offset >= len(iterator.runes) {
		iterator.error = errors.Fail(iter.ErrEndIteration{}, nil, "No more tokens")
	}

	return TextToken{value, style}, nil
}