package robotname

// Robot is a robot struct
type Robot struct {
	name string
}

// Name returns the name of the robot
func (r Robot) Name() (string, error) {
	return r.name, nil
}

// Reset clears the robot's name
func (r Robot) Reset() {
	r.name = ""
}
