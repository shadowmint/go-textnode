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
			Style: v.Style}
		rtn.Styles = make(map[string]string)
		for k, v := range text.Styles {
			rtn.Styles[string(k)] = v
		}
	}
	return &rtn, nil
}

func (t *TextTemplate) AsNode() *TextNode {
	rtn := NewTextNode()
	for k, v := range t.Nodes {
		if len(v.Value) > 0 {
			rtn.Text(k, v.Value)
		}
		if len(v.Style) > 0 {
			rtn.Style(k, v.Style)
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