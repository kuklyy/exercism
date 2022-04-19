// Package weather provides a weather forecast for certain city.
package weather

// CurrentCondition represents a condition for certain city.
var CurrentCondition string

// CurrentLocation represents a city for forecast.
var CurrentLocation string

// Forecast returns a weather forecast for certain city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
