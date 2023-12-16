package purchase

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	return kind == "car" || kind == "truck"
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	out := option1
	if option2 < option1 {
		out = option2
	}
	return out + " is clearly the better choice."
}

// 3 years old, it costs 80% of the original price it had when it was brand new.
// If it is at least 10 years old, it costs 50%.
// If the vehicle is at least 3 years old but less than 10 years, it costs 70% of the original pric

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	if age <= 3 {
		return originalPrice * .8
	} else if age >= 10 {
		return originalPrice * .5
	} else {
		return originalPrice * .7
	}
}
