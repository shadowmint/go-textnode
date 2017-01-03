package textnode

import (
	"golang.org/x/text/language"
	"fmt"
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
	Values      map[language.Tag]string
	Styles      map[language.Tag]string
	Constraints map[string]Constraint
}

// NewTextNode returns a new TextNode
func NewTextNode() *TextNode {
	return &TextNode{
		nodes: make(map[string]*textNodeEntry),
		Styles: make(map[rune]string)}
}

// Text sets a text entry for this node
func (n *TextNode) Text(locale string, id string, value string) error {
	node := n.getNode(id)

	tag := language.Make(locale)
	_, isDup := node.Values[tag]
	if isDup {
		return errors.Fail(ErrDuplicateId{}, nil, fmt.Sprintf("Duplicate locale / id match: %s / %s", locale, id))
	}

	node.Values[tag] = value

	// Add missing style if missing
	if _, ok := node.Styles[tag]; !ok {
		node.Styles[tag] = strings.Repeat(" ", utf8.RuneCountInString(value))
	}

	return nil
}

// Style sets the styles for a text entry for this node
func (n *TextNode) Style(locale string, id string, value string) error {
	node := n.getNode(id)
	tag := language.Make(locale)
	node.Styles[tag] = value

	// Add missing value if missing
	if _, ok := node.Values[tag]; !ok {
		node.Values[tag] = strings.Repeat(" ", utf8.RuneCountInString(value))
	}

	return nil
}

// Constraint adds a constraint for a given text entry
func (n *TextNode) Constraint(id string, statusId string, statusType int, threshold float32) {
	node := n.getNode(id)
	node.Constraints[statusId] = Constraint{
		Type: statusType,
		Threshold: threshold}
}

// Resolve the text representation of this node
func (n *TextNode) Resolve(env *Env) (*Text, error) {
	// Find the first target that matches the given env
	var target *textNodeEntry = nil
	for _, v := range n.nodes {
		matches := true
		for id, constraint := range v.Constraints {
			if !constraint.Meets(id, env) {
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

	// Find the language node to return
	for tag := range target.Values {
		best, _, _ := env.Matcher.Match(tag)
		if env.Language != nil && best == *env.Language {
			return newText(target.Values[tag], target.Styles[tag], n.Styles, env), nil
		}
	}

	return nil, errors.Fail(ErrNoText{}, nil, "No text entry for the selected language")
}

// getNode returns the entry record for the given id
func (n *TextNode) getNode(id string) *textNodeEntry {
	node, found := n.nodes[id]
	if !found {
		n.nodes[id] = &textNodeEntry{
			Values: make(map[language.Tag]string),
			Styles: make(map[language.Tag]string),
			Constraints: make(map[string]Constraint)}
		node = n.nodes[id]
	}
	return node
}