package textnode_test

import (
	"ntoolkit/assert"
	"ntoolkit/textnode"
	"testing"
	"golang.org/x/text/language"
)

func TestTemplate(T *testing.T) {
	assert.Test(T, func(T *assert.T) {
		node := textnode.NewTextNode()

		node.Text("en", "description", "English light")
		node.Text("de", "description", "German light")
		node.Constraint("description", "light", textnode.GreaterThanEq, 0.5)

		node.Text("en", "description.dark", "English dark")
		node.Text("de", "description.dark", "German dark")
		node.Constraint("description.dark", "light", textnode.LessThan, 0.5)

		template, err := textnode.TextTemplateFromTextNode(node)
		T.Assert(err == nil)
		T.Assert(template != nil)

		node2 := template.AsNode()
		T.Assert(node2 != nil)

		status := textnode.NewStatus()
		envBase := textnode.NewEnv(status, language.AmericanEnglish, language.German)
		envEN := envBase.SelectLanguage("en")
		envDE := envBase.SelectLanguage("de")

		status.Values["light"] = 1.0
		en1, err1 := node2.Resolve(envEN)
		de1, err2 := node2.Resolve(envDE)

		status.Values["light"] = 0.0
		en2, err3 := node2.Resolve(envEN)
		de2, err4 := node2.Resolve(envDE)

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

func TestTemplateAsJson(T *testing.T) {
	assert.Test(T, func(T *assert.T) {
		node := textnode.NewTextNode()

		node.Text("en", "description", "English light")
		node.Text("de", "description", "German light")
		node.Constraint("description", "light", textnode.GreaterThanEq, 0.5)

		node.Text("en", "description.dark", "English dark")
		node.Text("de", "description.dark", "German dark")
		node.Constraint("description.dark", "light", textnode.LessThan, 0.5)

		node.Styles[' '] = "Normal"

		template, err := textnode.TextTemplateFromTextNode(node)
		T.Assert(err == nil)
		T.Assert(template != nil)

		json, err := template.AsJson()
		T.Assert(err == nil)
		T.Assert(json != "")
	})
}

func TestTemplateFromJson(T *testing.T) {
	assert.Test(T, func(T *assert.T) {
		json := `{
			"Nodes": {
				"description": {
					"Values": {
						"de": "German light",
						"en": "English light"
					},
					"Styles": {
						"de": "            ",
						"en": "             "
					},
					"Constraints": {
						"light": {
							"Id": "light",
							"Type": "\u003e=",
							"Threshold": 0.5
						}
					}
				},
				"description.dark": {
					"Values": {
						"de": "German dark",
						"en": "English dark"
					},
					"Styles": {
						"de": "           ",
						"en": "            "
					},
					"Constraints": {
						"light": {
							"Id": "light",
							"Type": "\u003c",
							"Threshold": 0.5
						}
					}
				}
			},
			"Styles": {
				" ": "Normal"
			}
		}`;

		template, err := textnode.TextTemplateFromJson(json)
		T.Assert(err == nil)
		T.Assert(template != nil)

		node2 := template.AsNode()
		T.Assert(node2 != nil)

		status := textnode.NewStatus()
		envBase := textnode.NewEnv(status, language.AmericanEnglish, language.German)
		envEN := envBase.SelectLanguage("en")
		envDE := envBase.SelectLanguage("de")

		status.Values["light"] = 1.0
		en1, err1 := node2.Resolve(envEN)
		de1, err2 := node2.Resolve(envDE)

		status.Values["light"] = 0.0
		en2, err3 := node2.Resolve(envEN)
		de2, err4 := node2.Resolve(envDE)

		T.Assert(err1 == nil)
		T.Assert(err2 == nil)
		T.Assert(err3 == nil)
		T.Assert(err4 == nil)

		T.Assert(en1.Value == "English light")
		T.Assert(de1.Value == "German light")

		T.Assert(en2.Value == "English dark")
		T.Assert(de2.Value == "German dark")
	});
}