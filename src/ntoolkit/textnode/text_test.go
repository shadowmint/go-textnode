package textnode_test

import (
	"testing"
	"ntoolkit/assert"
	"ntoolkit/textnode"
)

func TestNew(T *testing.T) {
	assert.Test(T, func(T *assert.T) {
		T.Assert(textnode.New() != nil)
	})
}

func TestResolveMatchingSymbol(T *testing.T) {
	assert.Test(T, func(T *assert.T) {
		txt := textnode.New()

		txt.Import(map[string]string{
			"description": "description value goes here",
			"other":       "other value *here* yeah?",
			"empty":       "    "})

		renderer := &textnode.PassThroughRenderer{}

		output1, err := txt.Resolve("description", renderer)
		T.Assert(err == nil)
		T.Assert(output1 == "description value goes here")

		output2, err := txt.Resolve("other", renderer)
		T.Assert(err == nil)
		T.Assert(output2 == "other value *here* yeah?")

		output3, err := txt.Resolve("empty", renderer)
		T.Assert(err == nil)
		T.Assert(output3 == "")
	})
}

func TestResolveMissingSymbol(T *testing.T) {
	assert.Test(T, func(T *assert.T) {
		txt := textnode.New()

		txt.Import(map[string]string{
			"description": "description value goes here",
			"other":       "other value *here* yeah?",
			"empty":       "    "})

		renderer := &textnode.PassThroughRenderer{}

		output1, err := txt.Resolve("description22", renderer)
		T.Assert(err != nil)
		T.Assert(output1 == "")

		txt = textnode.New()

		output1, err = txt.Resolve("description22", renderer)
		T.Assert(err != nil)
		T.Assert(output1 == "")
	})
}