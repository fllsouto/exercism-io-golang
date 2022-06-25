// Package weather responsible to make weather forecast.
package weather

// CurrentCondition variable that store weather condition.
var CurrentCondition string

// CurrentLocation variable that store city localization.
var CurrentLocation string

// Forecast weather conditions for a given city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
