package textnode_test

import (
	"ntoolkit/assert"
	"ntoolkit/textnode"
	"testing"
)

func TestTemplate(T *testing.T) {
	assert.Test(T, func(T *assert.T) {
		node := textnode.NewTextNode()

		node.Text("description.light", "light")
		node.Text("description.dark", "dark")

		template, err := textnode.TextTemplateFromTextNode(node)
		T.Assert(err == nil)
		T.Assert(template != nil)

		node2 := template.AsNode()
		T.Assert(node2 != nil)

		styles := textnode.NewStyleSheet()

		en1, err1 := node2.Resolve(styles, "description.light")
		en2, err2 := node2.Resolve(styles, "description.dark")

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
		node.Text("description.dark", "dark")

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
				"description.light": {
					"Value": "light",
					"Style": "     "
				},
				"description.dark": {
					"Value": "dark",
					"Style": "    "
				}
			},
			"Styles": {
				" ": "Normal"
			}
		}`

		template, err := textnode.TextTemplateFromJson(json)
		T.Assert(err == nil)
		T.Assert(template != nil)

		node2 := template.AsNode()
		T.Assert(node2 != nil)

		styles := textnode.NewStyleSheet()

		en1, err1 := node2.Resolve(styles, "description.light")
		en2, err2 := node2.Resolve(styles, "description.dark")

		T.Assert(err1 == nil)
		T.Assert(err2 == nil)

		T.Assert(en1.Value == "light")
		T.Assert(en2.Value == "dark")
	})
}

func TestPartialTemplateFromJson(T *testing.T) {
	assert.Test(T, func(T *assert.T) {
		json := `{
			"Nodes": {
				"description.light": {
					"Value": "light"
				},
				"description.dark": {
					"Style": "    "
				}
			},
			"Styles": {
				" ": "Normal"
			}
		}`

		template, err := textnode.TextTemplateFromJson(json)
		T.Assert(err == nil)
		T.Assert(template != nil)

		node2 := template.AsNode()
		T.Assert(node2 != nil)

		styles := textnode.NewStyleSheet()

		en1, err1 := node2.Resolve(styles, "description.light")
		en2, err2 := node2.Resolve(styles, "description.dark")

		T.Assert(err1 == nil)
		T.Assert(err2 == nil)

		T.Assert(en1.Value == "light")
		T.Assert(en2.Value == "    ")
	})
}
