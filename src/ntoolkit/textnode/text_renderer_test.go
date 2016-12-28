package textnode_test

import (
	"ntoolkit/textnode"
	"testing"
	"ntoolkit/assert"
	"golang.org/x/text/language"
	"fmt"
)

type FakeRenderer struct {
}

func (r *FakeRenderer) AsString(t textnode.TextToken) (string, error) {
	rtn := t.Value
	if t.Style.Decoration & textnode.Bold != 0 {
		rtn = fmt.Sprintf("*%s*", rtn)
	}
	if t.Style.Decoration & textnode.Underline != 0 {
		rtn = fmt.Sprintf("__%s__", rtn)
	}
	return rtn, nil
}

type TemplateProps struct {
	Monster string
	Weapon  string
	Player  string
	Damage  int
}

func TestRenderText(T *testing.T) {
	assert.Test(T, func(T *assert.T) {
		node := textnode.NewTextNode()
		node.Text("en", "description", "The {{.Monster}}'s {{.Weapon}} smashed into {{.Player}}'s leg, dealing {{.Damage}} damage.")
		node.Style("en", "description", "    uuuuuuuuuuuuuu                          uuuuuuuuuuuuu      bbbbbbbbbbbbbbbbbbbbbbbbbb ")
		node.Styles['u'] = "underline"
		node.Styles['b'] = "bold"

		status := textnode.NewStatus()
		env := textnode.NewEnv(status, language.AmericanEnglish).SelectLanguage("en")

		s1 := env.Stylesheet.New("underline")
		s1.Decoration |= textnode.Underline

		s2 := env.Stylesheet.New("bold")
		//s2.Decoration |= textnode.Underline
		s2.Decoration |= textnode.Bold

		en1, err1 := node.Resolve(env)

		T.Assert(len(en1.Errors) == 0)
		T.Assert(err1 == nil)
		T.Assert(en1 != nil)

		props := TemplateProps{
			Monster: "Troll",
			Weapon: "Nerf-Hammer",
			Player: "Elric",
			Damage: 0}
		output, err2 := en1.RenderTemplate(props, &FakeRenderer{})

		T.Assert(err2 == nil)
		T.Assert(output == "The __Troll's__ Nerf-Hammer smashed into __Elric's__ leg, *dealing 0 damage*.")

		props2 := TemplateProps{
			Monster: "Troll",
			Weapon: "Slap",
			Player: "Elric",
			Damage: 100}
		output2, err3 := en1.RenderTemplate(props2, &FakeRenderer{})

		T.Assert(err3 == nil)
		T.Assert(output2 == "The __Troll's__ Slap smashed into __Elric's__ leg, *dealing 100 damage*.")
	})
}