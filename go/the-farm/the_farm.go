package thefarm

import (
	"errors"
	"fmt"
)

func DivideFood(fc FodderCalculator, numberOfCows int) (float64, error) {
	fatFactor, err := fc.FatteningFactor()
	if err != nil {
		return 0, err
	}
	totalFodder, err := fc.FodderAmount(numberOfCows)
	if err != nil {
		return 0, err
	}
	totalFodder *= fatFactor
	return totalFodder / float64(numberOfCows), nil
}

func ValidateInputAndDivideFood(fc FodderCalculator, numberOfCows int) (float64, error) {
	if numberOfCows <= 0 {
		return 0, errors.New("invalid number of cows")
	}
	return DivideFood(fc, numberOfCows)
}

type InvalidCowsError struct {
	Message      string
	NumberOfCows int
}

func (e InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", e.NumberOfCows, e.Message)
}

func ValidateNumberOfCows(numberOfCows int) error {
	if numberOfCows < 0 {
		return &InvalidCowsError{Message: "there are no negative cows", NumberOfCows: numberOfCows}
	}
	if numberOfCows == 0 {
		return &InvalidCowsError{Message: "no cows don't need food", NumberOfCows: numberOfCows}
	}
	return nil
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
