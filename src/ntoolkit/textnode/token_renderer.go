package textnode

type TokenRenderer interface {
	// AsString should return the appropriate device formatted string for the given token value.
	AsString(text TextToken) (string, error)
}