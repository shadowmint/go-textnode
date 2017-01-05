package textnode

import (
	"encoding/json"
	"strings"
	"bytes"
)

type TextTemplate struct {
	Nodes  map[string]TextTemplateEntry
	Styles map[string]string
}

type TextTemplateEntry struct {
	Value       string
	Style       string
	Constraints map[string]TextConstraintTemplate
}

type TextConstraintTemplate struct {
	Type      string
	Threshold float32
}

func TextTemplateFromJson(raw string) (*TextTemplate, error) {
	rtn := TextTemplate{}
	dec := json.NewDecoder(strings.NewReader(raw))
	err := dec.Decode(&rtn)
	if err != nil {
		return nil, err
	}
	return &rtn, nil
}

func TextTemplateFromTextNode(text *TextNode) (*TextTemplate, error) {
	rtn := TextTemplate{Nodes: make(map[string]TextTemplateEntry)}
	for k, v := range text.nodes {
		rtn.Nodes[k] = TextTemplateEntry{
			Value: v.Value,
			Style: v.Style,
			Constraints: make(map[string]TextConstraintTemplate)}
		for nk, nv := range v.Constraints {
			rtn.Nodes[k].Constraints[nk] = TextConstraintTemplate{
				Type: constraintValueAsKey(nv.Type),
				Threshold: nv.Threshold}
		}
	}
	rtn.Styles = make(map[string]string)
	for k, v := range text.Styles {
		rtn.Styles[string(k)] = v
	}
	return &rtn, nil
}

func (t *TextTemplate) AsNode() *TextNode {
	rtn := NewTextNode()
	for k, v := range t.Nodes {
		rtn.Text(k, v.Value)
		rtn.Style(k, v.Style)
		for nk, nv := range v.Constraints {
			rtn.Constraint(k, nk, constraintKeyAsValue(nv.Type), nv.Threshold)
		}
	}
	for k, v := range t.Styles {
		values := []rune(k)
		if len(values) > 0 {
			rtn.Styles[values[0]] = v
		}
	}
	return rtn
}

func (t *TextTemplate) AsJson() (string, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return "", err
	}
	var out bytes.Buffer
	json.Indent(&out, b, "", "\t")
	rtn := string(out.Bytes())
	return rtn, err
}

func constraintKeyAsValue(operator string) int {
	var op int = Equal
	switch operator {
	case "==":
		op = Equal
		break
	case "!=":
		op = NotEqual
		break
	case ">":
		op = GreaterThan
		break
	case ">=":
		op = GreaterThanEq
		break
	case "<":
		op = LessThan
		break
	case "<=":
		op = LessThanEq
		break
	}
	return op
}

func constraintValueAsKey(operator int) string {
	var op string = "=="
	switch operator {
	case Equal:
		op = "=="
		break
	case NotEqual:
		op = "!="
		break
	case GreaterThan:
		op = ">"
		break
	case GreaterThanEq:
		op = ">="
		break
	case LessThan:
		op = "<"
		break
	case LessThanEq:
		op = "<="
		break
	}
	return op
}