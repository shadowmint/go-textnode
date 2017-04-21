package textnode

import "ntoolkit/parser"

type NoOpClassifier struct {
}

func (c *NoOpClassifier) Classify(token *parser.Token) bool {
	return false
}
