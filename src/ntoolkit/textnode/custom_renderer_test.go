package textnode_test

import (
	"ntoolkit/parser"
	"strings"
	"testing"
	"ntoolkit/assert"
	"ntoolkit/textnode"
)

const (
	CapsLower = 1
	CapsUpper = 2
)

// **foo** -> FOO, and *FoO* -> foo
type CapsClassifier struct {
}

func (c *CapsClassifier) Classify(token *parser.Token) bool {
	if strings.HasPrefix(*token.Raw, "**") && strings.HasSuffix(*token.Raw, "**") {
		if !token.Is(CapsUpper) {
			token.Type = CapsUpper
			value := strings.ToUpper((*token.Raw)[2:len(*token.Raw)-2])
			token.Raw = &value
			return true
		}
	} else if strings.HasPrefix(*token.Raw, "*") && strings.HasSuffix(*token.Raw, "*") {
		if !token.Is(CapsLower) {
			token.Type = CapsLower
			value := strings.ToLower((*token.Raw)[1:len(*token.Raw)-1])
			token.Raw = &value
			return true
		}
	}
	return false
}

type CapsRenderer struct {
}

func (r *CapsRenderer) Render(tokens *parser.Tokens) (string, error) {
	if tokens == nil {
		return "", nil
	}
	if tokens.Count() == 0 {
		return "", nil
	}
	return tokens.Front.WalkRaw(" "), nil
}

func TestCustomerRenderer(T *testing.T) {
	assert.Test(T, func(T *assert.T) {
		txt := textnode.New()

		txt.Import(map[string]string{"description": "description **value** goes here *ShouldBeLower*"})

		renderer := &CapsRenderer{}
		classifier := &CapsClassifier{}

		output1, err := txt.Resolve("description", renderer, classifier)
		T.Assert(err == nil)
		T.Assert(output1 == "description VALUE goes here shouldbelower")
	})
}
