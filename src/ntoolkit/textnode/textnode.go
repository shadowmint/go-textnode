package textnode

import (
	"ntoolkit/errors"
	"strings"
	"unicode/utf8"
)

// TextNode is the basic descriptive unit for an object, it contains various LODs for
// different environmental conditions.
type TextNode struct {
	nodes  map[string]*textNodeEntry
	Styles map[rune]string
}

// textNodeEntry is a single record in the text node.
type textNodeEntry struct {
	Value string
	Style string
	Constraints map[string]Constraint
}

// NewTextNode returns a new TextNode
func NewTextNode() *TextNode {
	return &TextNode{
		nodes: make(map[string]*textNodeEntry),
		Styles: make(map[rune]string)}
}

// Keys returns the set of defined keys for this node
func (n *TextNode) Keys() []string {
  rtn := make([]string, 0)
  for k, _ := range n.nodes {
    rtn = append(rtn, k)
  }
  return rtn
}

// Text sets a text entry for this node
func (n *TextNode) Text(id string, value string) error {
	node := n.getNode(id)
	node.Value = value
	if len(node.Style) != len(node.Value) {
		node.Style = strings.Repeat(" ", utf8.RuneCountInString(value))
	}
	return nil
}

// Style sets the styles for a text entry for this node
func (n *TextNode) Style(id string, value string) error {
	node := n.getNode(id)
	node.Style = value
	if len(node.Style) != len(node.Value) {
		node.Value = strings.Repeat(" ", utf8.RuneCountInString(value))
	}
	return nil
}

// Constraint adds a constraint for a given text entry
func (n *TextNode) Constraint(id string, statusId string, statusType int, threshold float32) {
	node := n.getNode(id)
	node.Constraints[statusId] = Constraint{
		Id: statusId,
		Type: statusType,
		Threshold: threshold}
}

// Resolve the text representation of this node
func (n *TextNode) Resolve(env *Env) (*Text, error) {
	// Find the first target that matches the given env
	var target *textNodeEntry = nil
	for _, v := range n.nodes {
		matches := true
		for _, constraint := range v.Constraints {
			if !constraint.Meets(env) {
				matches = false
				break
			}
		}
		if matches {
			target = v
			break
		}
	}
	if target == nil {
		return nil, errors.Fail(ErrNoText{}, nil, "No text entry for the given constraints")
	}

	return newText(target.Value, target.Style, n.Styles, env), nil
}

// getNode returns the entry record for the given id
func (n *TextNode) getNode(id string) *textNodeEntry {
	node, found := n.nodes[id]
	if !found {
		n.nodes[id] = &textNodeEntry{
			Value: "",
			Style: "",
			Constraints: make(map[string]Constraint)}
		node = n.nodes[id]
	}
	return node
}
