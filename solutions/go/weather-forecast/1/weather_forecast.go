// Package weather tells the current weather confiton in a certain location.
package weather

// CurrentCondition indicates the weather condition.
var CurrentCondition string
// CurrentLocation indicates the location that we want to prompt.
var CurrentLocation string

// Forecast takes a city and a condition as strings, concatenates them and the function returns the current weather condition of an indicated location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
