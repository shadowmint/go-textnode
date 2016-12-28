package textnode

// ErrDuplicateId is returned when an attempt is made to add a text node
// when the given id has already been used.
type ErrDuplicateId struct{}

// ErrNoText is returned when a text node cannot be resolved into a text value.
type ErrNoText struct{}

// ErrBadStyles is raised when an attempt is made to process a Text object with invalid style properties.
type ErrBadStyles struct{}