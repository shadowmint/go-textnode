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
		node.Constraint("description", "light", textnode.GreaterThanEq, 0.5)

		node.Text("description.dark", "dark")
		node.Constraint("description.dark", "light", textnode.LessThan, 0.5)

		status := textnode.NewStatus()
		env := textnode.NewEnv(status)

		status.Values["light"] = 1.0
		en1, err1 := node.Resolve(env)

		status.Values["light"] = 0.0
		en2, err2 := node.Resolve(env)

		T.Assert(err1 == nil)
		T.Assert(err2 == nil)

		T.Assert(en1.Value == "light")
		T.Assert(en2.Value == "dark")
	})
}

func TestStyledTextNode(T *testing.T) {
	assert.Test(T, func(T *assert.T) {
		node := textnode.NewTextNode()
		node.Text("description",  "Hello World")
		node.Style("description", "rrrrr bbbbb")
		node.Styles['r'] = "red"
		node.Styles['b'] = "blue"

		status := textnode.NewStatus()
		env := textnode.NewEnv(status)
		s1 := env.Stylesheet.New("red")
		s2 := env.Stylesheet.New("blue")

		en1, err1 := node.Resolve(env)

		T.Assert(err1 == nil)
		T.Assert(en1 != nil)

		T.Assert(en1.Value == "Hello World")
		T.Assert(en1.Styles[0] == s1)
		T.Assert(en1.Styles[1] == s1)
		T.Assert(en1.Styles[2] == s1)
		T.Assert(en1.Styles[3] == s1)
		T.Assert(en1.Styles[4] == s1)
		T.Assert(en1.Styles[5] == &env.Stylesheet.Default)
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
		node.Text("description",  "Hello World")
		node.Style("description", "rrrrr bbbbb")
		node.Styles['b'] = "blue"

		status := textnode.NewStatus()
		env := textnode.NewEnv(status)
		env.Stylesheet.New("red")
		s2 := env.Stylesheet.New("blue")

		en1, err1 := node.Resolve(env)

		T.Assert(err1 == nil)
		T.Assert(en1 != nil)

		T.Assert(errors.Is(en1.Errors[0], textnode.ErrBadStyles{}))
		T.Assert(en1.Value == "Hello World")
		T.Assert(en1.Styles[0] == &env.Stylesheet.Default)
		T.Assert(en1.Styles[1] == &env.Stylesheet.Default)
		T.Assert(en1.Styles[2] == &env.Stylesheet.Default)
		T.Assert(en1.Styles[3] == &env.Stylesheet.Default)
		T.Assert(en1.Styles[4] == &env.Stylesheet.Default)
		T.Assert(en1.Styles[5] == &env.Stylesheet.Default)
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
		node.Text("description",  "Hello World")
		node.Style("description", "rrrrr bbbbb")
		node.Styles['r'] = "red"
		node.Styles['b'] = "blue"

		status := textnode.NewStatus()
		env := textnode.NewEnv(status)
		s2 := env.Stylesheet.New("blue")

		en1, err1 := node.Resolve(env)

		T.Assert(err1 == nil)
		T.Assert(en1 != nil)

		T.Assert(errors.Is(en1.Errors[0], textnode.ErrBadStyles{}))
		T.Assert(en1.Value == "Hello World")
		T.Assert(en1.Styles[0] == &env.Stylesheet.Default)
		T.Assert(en1.Styles[1] == &env.Stylesheet.Default)
		T.Assert(en1.Styles[2] == &env.Stylesheet.Default)
		T.Assert(en1.Styles[3] == &env.Stylesheet.Default)
		T.Assert(en1.Styles[4] == &env.Stylesheet.Default)
		T.Assert(en1.Styles[5] == &env.Stylesheet.Default)
		T.Assert(en1.Styles[6] == s2)
		T.Assert(en1.Styles[7] == s2)
		T.Assert(en1.Styles[8] == s2)
		T.Assert(en1.Styles[9] == s2)
		T.Assert(en1.Styles[10] == s2)
	})
}
