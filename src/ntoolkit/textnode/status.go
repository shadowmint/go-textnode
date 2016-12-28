package textnode

// Status is a shared set of status properties for a specific local area.
type Status struct {
	Values map[string]float32
}

// NewStatus returns a new status object pointer
func NewStatus() *Status {
	return &Status{
		Values: make(map[string]float32)}
}
