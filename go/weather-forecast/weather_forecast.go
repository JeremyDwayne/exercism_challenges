// Package weather provides weather forecast functionality.
package weather

// CurrentCondition is a variable that holds the current weather condition.
var CurrentCondition string

// CurrentLocation is a variable that holds the current location.
var CurrentLocation string

// Forecast returns a weather forecast for a given city and condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
