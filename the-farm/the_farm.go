package thefarm

import (
	"errors"
	"fmt"
)

func DivideFood(fodderCalculator FodderCalculator, numberOfCows int) (float64, error) {
	fodderAmount, err := fodderCalculator.FodderAmount(numberOfCows)

	if err != nil {
		return 0, err
	}

	fatteningFactor, err := fodderCalculator.FatteningFactor()

	if err != nil {
		return 0, err
	}

	return (fodderAmount * fatteningFactor) / float64(numberOfCows), nil
}

func ValidateInputAndDivideFood(fodderCalculator FodderCalculator, numberOfCows int) (float64, error) {
	if numberOfCows <= 0 {
		return 0, errors.New("invalid number of cows")
	}

	return DivideFood(fodderCalculator, numberOfCows)
}

type InvalidCowsError struct {
	numberOfCows int
	message      string
}

func (err *InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", err.numberOfCows, err.message)
}

func ValidateNumberOfCows(numberOfCows int) error {
	if numberOfCows == 0 {
		return &InvalidCowsError{numberOfCows: numberOfCows, message: "no cows don't need food"}
	}

	if numberOfCows < 0 {
		return &InvalidCowsError{numberOfCows: numberOfCows, message: "there are no negative cows"}
	}

	return nil
}
