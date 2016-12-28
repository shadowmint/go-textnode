package textnode_test

import (
	"ntoolkit/assert"
	"ntoolkit/textnode"
	"testing"
	"golang.org/x/text/language"
	"ntoolkit/errors"
)

func TestSimpleTextNode(T *testing.T) {
	assert.Test(T, func(T *assert.T) {
		node := textnode.NewTextNode()

		node.Text("en", "description", "English light")
		node.Text("de", "description", "German light")
		node.Constraint("description", "light", textnode.GreaterThanEq, 0.5)

		node.Text("en", "description.dark", "English dark")
		node.Text("de", "description.dark", "German dark")
		node.Constraint("description.dark", "light", textnode.LessThan, 0.5)

		status := textnode.NewStatus()
		envBase := textnode.NewEnv(status, language.AmericanEnglish, language.German)
		envEN := envBase.SelectLanguage("en")
		envDE := envBase.SelectLanguage("de")

		status.Values["light"] = 1.0
		en1, err1 := node.Resolve(envEN)
		de1, err2 := node.Resolve(envDE)

		status.Values["light"] = 0.0
		en2, err3 := node.Resolve(envEN)
		de2, err4 := node.Resolve(envDE)

		T.Assert(err1 == nil)
		T.Assert(err2 == nil)
		T.Assert(err3 == nil)
		T.Assert(err4 == nil)

		T.Assert(en1.Value == "English light")
		T.Assert(de1.Value == "German light")

		T.Assert(en2.Value == "English dark")
		T.Assert(de2.Value == "German dark")
	})
}

func TestStyledTextNode(T *testing.T) {
	assert.Test(T, func(T *assert.T) {
		node := textnode.NewTextNode()
		node.Text("en", "description",  "Hello World")
		node.Style("en", "description", "rrrrr bbbbb")
		node.Styles['r'] = "red"
		node.Styles['b'] = "blue"

		status := textnode.NewStatus()
		env := textnode.NewEnv(status, language.AmericanEnglish).SelectLanguage("en")
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
		node.Text("en", "description",  "Hello World")
		node.Style("en", "description", "rrrrr bbbbb")
		node.Styles['b'] = "blue"

		status := textnode.NewStatus()
		env := textnode.NewEnv(status, language.AmericanEnglish).SelectLanguage("en")
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
		node.Text("en", "description",  "Hello World")
		node.Style("en", "description", "rrrrr bbbbb")
		node.Styles['r'] = "red"
		node.Styles['b'] = "blue"

		status := textnode.NewStatus()
		env := textnode.NewEnv(status, language.AmericanEnglish).SelectLanguage("en")
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
