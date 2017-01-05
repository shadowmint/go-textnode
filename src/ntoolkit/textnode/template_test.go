package textnode_test

import (
	"ntoolkit/assert"
	"ntoolkit/textnode"
	"testing"
)

func TestTemplate(T *testing.T) {
	assert.Test(T, func(T *assert.T) {
		node := textnode.NewTextNode()

		node.Text("description", "light")
		node.Constraint("description", "light", textnode.GreaterThanEq, 0.5)

		node.Text("description.dark", "dark")
		node.Constraint("description.dark", "light", textnode.LessThan, 0.5)

		template, err := textnode.TextTemplateFromTextNode(node)
		T.Assert(err == nil)
		T.Assert(template != nil)

		node2 := template.AsNode()
		T.Assert(node2 != nil)

		status := textnode.NewStatus()
		env := textnode.NewEnv(status)

		status.Values["light"] = 1.0
		en1, err1 := node2.Resolve(env)

		status.Values["light"] = 0.0
		en2, err2 := node2.Resolve(env)

		T.Assert(err1 == nil)
		T.Assert(err2 == nil)

		T.Assert(en1.Value == "light")
		T.Assert(en2.Value == "dark")
	})
}

func TestTemplateAsJson(T *testing.T) {
	assert.Test(T, func(T *assert.T) {
		node := textnode.NewTextNode()

		node.Text("description", "light")
		node.Constraint("description", "light", textnode.GreaterThanEq, 0.5)

		node.Text("description.dark", "dark")
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
					"Value": "light",
					"Style": "     ",
					"Constraints": {
						"light": {
							"Type": "\u003e=",
							"Threshold": 0.5
						}
					}
				},
				"description.dark": {
					"Value": "dark",
					"Style": "    ",
					"Constraints": {
						"light": {
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
		env := textnode.NewEnv(status)

		status.Values["light"] = 1.0
		en1, err1 := node2.Resolve(env)

		status.Values["light"] = 0.0
		en2, err2 := node2.Resolve(env)

		T.Assert(err1 == nil)
		T.Assert(err2 == nil)

		T.Assert(en1.Value == "light")
		T.Assert(en2.Value == "dark")
	});
}