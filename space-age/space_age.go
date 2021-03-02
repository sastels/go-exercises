package space

// Planet is the name of a planet
type Planet string

var years = map[Planet]float64{
	"Earth":   31557600.0,
	"Mercury": 0.2408467 * 31557600.0,
	"Venus":   0.61519726 * 31557600.0,
	"Mars":    1.8808158 * 31557600.0,
	"Jupiter": 11.862615 * 31557600.0,
	"Saturn":  29.447498 * 31557600.0,
	"Uranus":  84.016846 * 31557600.0,
	"Neptune": 164.79132 * 31557600.0,
}

// Age converts the age in seconds to age in planet-years
func Age(seconds float64, planet Planet) float64 {
	return seconds / years[planet]
}
