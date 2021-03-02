// Package gigasecond adds gigaseconds to things
package gigasecond

import "time"

// Add a gigasecond to t
func AddGigasecond(t time.Time) time.Time {
	d, _ := time.ParseDuration("1000000000s")
	return t.Add(d)
}
