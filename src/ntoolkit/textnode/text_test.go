package textnode_test

import (
	"ntoolkit/assert"
	"ntoolkit/textnode"
	"testing"
	"ntoolkit/iter"
)

func TestTextTokens(T *testing.T) {
	assert.Test(T, func(T *assert.T) {
		sheet := textnode.NewStyleSheet()
		style1 := sheet.New("red")
		style2 := sheet.New("blue")

		text := textnode.Text{
			Value: "Hello 正解!",
			Styles: []*textnode.Style{style1, style1, style1, style1, style1, style1, style2, style2, style1}}

		it, err := iter.Collect(text.Tokens())
		T.Assert(err == nil)

		T.Assert(len(it) == 3)
		T.Assert(it[0].(textnode.TextToken).Value == "Hello ")
		T.Assert(it[0].(textnode.TextToken).Style == style1)
		T.Assert(it[1].(textnode.TextToken).Value == "正解")
		T.Assert(it[1].(textnode.TextToken).Style == style2)
		T.Assert(it[2].(textnode.TextToken).Value == "!")
		T.Assert(it[2].(textnode.TextToken).Style == style1)
	})
}
