package clock

import "fmt"

// Clock represents a clock
type Clock struct {
	hour, min int
}

// New returns a new clock
func New(hour, min int) Clock {
	if min < 0 {
		hour -= 1 + (-min / 60)
		min += 60 + 60*(-min/60)
	}
	hour += (min / 60)
	min = min % 60
	if hour < 0 {
		hour += 24 + 24*(-hour/24)
	}
	hour = hour % 24
	return Clock{hour, min}
}

// Add adds minutes to a clock
func (c Clock) Add(min int) Clock {
	return New(c.hour, c.min+min)

}

// Subtract subtracts minutes from a clock
func (c Clock) Subtract(min int) Clock {
	return New(c.hour, c.min-min)
}

// String returns string representation of a clock
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.min)

}
