package textnode

import (
	"ntoolkit/errors"
	"fmt"
	"ntoolkit/parser"
	"sync"
)

type Text struct {
	literals  map[string]string // Literals is the set of normal key nodes
	tokenizer parser.Tokenizer
}

// New returns a new Text instance
func New() *Text {
	return &Text{literals: make(map[string]string)}
}

// Import takes a map and imports the raw values into the text object.
// Newly imported values take precedence.
func (t *Text) Import(values map[string]string) {
	for key := range values {
		t.literals[key] = values[key]
	}
}

// Export returns the internal raw string mappings
func (t *Text) Export() map[string]string {
  return t.literals
}

// Resolve returns the renderer string for the given key.
// If a classifier is provided it is used.
func (t *Text) Resolve(name string, renderer Renderer, classifier ...parser.Classifier) (string, error) {
	value, ok := t.literals[name]
	if !ok {
		return "", errors.Fail(ErrBadName{}, nil, fmt.Sprintf("No match for %s in text", name))
	}

	var classic parser.Classifier
	if len(classifier) > 0 {
		classic = classifier[0]
	} else {
		classic = &NoOpClassifier{}
	}

	var tokens *parser.Tokens
	var rerr error
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go (func() {
		parser.Parse(t.getTokenizer(), classic, func(handler func(data string, finished bool)) {
			handler(value, true)
		}).Then(func(v *parser.Tokens) {
			tokens = v
			wg.Done()
		}, func(e error) {
			rerr = e
			wg.Done()
		})
	})()
	wg.Wait()
	if rerr != nil {
		return "", rerr
	}

	return renderer.Render(tokens)
}

func (t *Text) getTokenizer() parser.Tokenizer {
	if t.tokenizer == nil {
		t.tokenizer = &wordTokenizer{}
	}
	return t.tokenizer
}
