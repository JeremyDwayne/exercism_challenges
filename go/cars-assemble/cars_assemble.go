package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	successRate /= 100
	return float64(productionRate) * successRate
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	productionRate /= 60
	successRate /= 100
	return int((float64(productionRate) * (successRate)))
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	var ungroupedCost uint = 10000
	var groupCost uint = 95000

	grouped := carsCount / 10
	ungrouped := carsCount % 10
	return (uint(grouped) * groupCost) + (uint(ungrouped) * ungroupedCost)
}
