package textnode

const (
	Equal = iota
	NotEqual = iota
	LessThan = iota
	LessThanEq = iota
	GreaterThan = iota
	GreaterThanEq = iota
)

// Env is represents the state for rendering for a particular user.
// eg. A user may have some status filter that renders different text.
type Env struct {
	Status     *Status
	Stylesheet *StyleSheet
}

// NewEnv returns a new Env with a matcher for the given set of supported languages.
func NewEnv(status *Status) *Env {
	return &Env{
		Status: status,
		Stylesheet: NewStyleSheet()}
}
