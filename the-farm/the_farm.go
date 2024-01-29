package thefarm

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
	panic("")
}

func ValidateNumberOfCows(numberOfCows int) error {
	panic("")
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
