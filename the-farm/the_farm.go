package thefarm

import "errors"

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

func ValidateNumberOfCows(numberOfCows int) error {
	panic("")
}
