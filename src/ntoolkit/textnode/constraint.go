package textnode

// Constraint in a condition to use to resolve what text node to render for an object.
type Constraint struct {
	Type      int
	Threshold float32
}

// Meets checks if the constraint meets the given env restrictions
func (c *Constraint) Meets(id string, env *Env) bool {
	value, ok := env.Status.Values[id]
	rtn := false
	if ok {
		switch c.Type {
		case Equal:
			rtn = value == c.Threshold
		case NotEqual:
			rtn = value != c.Threshold
		case GreaterThan:
			rtn = value > c.Threshold
		case GreaterThanEq :
			rtn = value >= c.Threshold
		case LessThan:
			rtn = value < c.Threshold
		case LessThanEq:
			rtn = value <= c.Threshold
		}
	}
	return rtn
}