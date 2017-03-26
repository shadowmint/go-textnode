package textnode_test

import (
	"ntoolkit/assert"
	"ntoolkit/textnode"
	"testing"
	"ntoolkit/errors"
)

func TestSimpleTextNode(T *testing.T) {
	assert.Test(T, func(T *assert.T) {
		node := textnode.NewTextNode()

		node.Text("description", "light")
		node.Text("description.dark", "dark")

		styles := textnode.NewStyleSheet()
		en1, err1 := node.Resolve(styles, "description")
		en2, err2 := node.Resolve(styles, "description.dark")

		T.Assert(err1 == nil)
		T.Assert(err2 == nil)

		T.Assert(en1.Value == "light")
		T.Assert(en2.Value == "dark")
	})
}

func TestStyledTextNode(T *testing.T) {
	assert.Test(T, func(T *assert.T) {
		node := textnode.NewTextNode()
		node.Text("description", "Hello World")
		node.Style("description", "rrrrr bbbbb")
		node.Styles['r'] = "red"
		node.Styles['b'] = "blue"

		styles := textnode.NewStyleSheet()
		s1 := styles.New("red")
		s2 := styles.New("blue")

		en1, err1 := node.Resolve(styles)

		T.Assert(err1 == nil)
		T.Assert(en1 != nil)

		T.Assert(en1.Value == "Hello World")
		T.Assert(en1.Styles[0] == s1)
		T.Assert(en1.Styles[1] == s1)
		T.Assert(en1.Styles[2] == s1)
		T.Assert(en1.Styles[3] == s1)
		T.Assert(en1.Styles[4] == s1)
		T.Assert(en1.Styles[5] == &styles.Default)
		T.Assert(en1.Styles[6] == s2)
		T.Assert(en1.Styles[7] == s2)
		T.Assert(en1.Styles[8] == s2)
		T.Assert(en1.Styles[9] == s2)
		T.Assert(en1.Styles[10] == s2)
	})
}

func TestStyledTextNodeMissingKey(T *testing.T) {
	assert.Test(T, func(T *assert.T) {
		node := textnode.NewTextNode()
		node.Text("description", "Hello World")
		node.Style("description", "rrrrr bbbbb")
		node.Styles['b'] = "blue"

		styles := textnode.NewStyleSheet()
		styles.New("red")
		s2 := styles.New("blue")

		en1, err1 := node.Resolve(styles)

		T.Assert(err1 == nil)
		T.Assert(en1 != nil)

		T.Assert(errors.Is(en1.Errors[0], textnode.ErrBadStyles{}))
		T.Assert(en1.Value == "Hello World")
		T.Assert(en1.Styles[0] == &styles.Default)
		T.Assert(en1.Styles[1] == &styles.Default)
		T.Assert(en1.Styles[2] == &styles.Default)
		T.Assert(en1.Styles[3] == &styles.Default)
		T.Assert(en1.Styles[4] == &styles.Default)
		T.Assert(en1.Styles[5] == &styles.Default)
		T.Assert(en1.Styles[6] == s2)
		T.Assert(en1.Styles[7] == s2)
		T.Assert(en1.Styles[8] == s2)
		T.Assert(en1.Styles[9] == s2)
		T.Assert(en1.Styles[10] == s2)
	})
}

func TestStyledTextNodeMissingStyle(T *testing.T) {
	assert.Test(T, func(T *assert.T) {
		node := textnode.NewTextNode()
		node.Text("description", "Hello World")
		node.Style("description", "rrrrr bbbbb")
		node.Styles['r'] = "red"
		node.Styles['b'] = "blue"

		styles := textnode.NewStyleSheet()
		s2 := styles.New("blue")

		en1, err1 := node.Resolve(styles)

		T.Assert(err1 == nil)
		T.Assert(en1 != nil)

		T.Assert(errors.Is(en1.Errors[0], textnode.ErrBadStyles{}))
		T.Assert(en1.Value == "Hello World")
		T.Assert(en1.Styles[0] == &styles.Default)
		T.Assert(en1.Styles[1] == &styles.Default)
		T.Assert(en1.Styles[2] == &styles.Default)
		T.Assert(en1.Styles[3] == &styles.Default)
		T.Assert(en1.Styles[4] == &styles.Default)
		T.Assert(en1.Styles[5] == &styles.Default)
		T.Assert(en1.Styles[6] == s2)
		T.Assert(en1.Styles[7] == s2)
		T.Assert(en1.Styles[8] == s2)
		T.Assert(en1.Styles[9] == s2)
		T.Assert(en1.Styles[10] == s2)
	})
}

func TestKeys(T *testing.T) {
	assert.Test(T, func(T *assert.T) {
		node := textnode.NewTextNode()
		node.Text("description1", "Hello World")
		node.Text("description2", "Hello World")
		node.Text("description3", "Hello World")
		node.Text("description4", "Hello World")
		node.Text("description5", "Hello World")
		keys := node.Keys()

		T.Assert(len(keys) == 5)
	})
}
