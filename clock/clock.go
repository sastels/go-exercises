package clock

// Clock represents a clock
type Clock struct {
	hour, min int
}

// New returns a new clock
func New(hour, min int) Clock {
	return Clock{hour, min}
}

// Add adds minutes to a clock
func (c Clock) Add(min int) Clock {
	return c
}

// Subtract subtracts minutes from a clock
func (c Clock) Subtract(min int) Clock {
	return c
}

func (c Clock) String() string {
	return ""
}
