package textnode

import (
	"fmt"
	"ntoolkit/iter"
	"ntoolkit/errors"
)

// TextTokenIter is an iterator that returns blocks of text
type TextTokenIter struct {
	error  error
	offset int
	length int
	text   *Text
	runes  []rune
}

func newTextTokenIter(text *Text) *TextTokenIter {
	runes := []rune(text.Value)
	return &TextTokenIter{
		error: nil,
		runes: runes,
		length: len(runes),
		offset: 0,
		text:text}
}

func (iterator *TextTokenIter) Next() (interface{}, error) {
	if iterator.error != nil {
		return nil, iterator.error
	}

	// Bounds check
	if len(iterator.text.Styles) != len(iterator.runes) {
		iterator.error = errors.Fail(ErrBadStyles{}, nil, fmt.Sprintf("Incorrect number of styles (%d) for string of length %d", len(iterator.text.Styles), len(iterator.runes)))
		return nil, iterator.error
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