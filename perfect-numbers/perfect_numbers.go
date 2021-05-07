package perfect

import "errors"

type Classification string

const ClassificationDeficient Classification = "ClassificationDeficient"
const ClassificationPerfect Classification = "ClassificationPerfect"
const ClassificationAbundant Classification = "ClassificationAbundant"

var ErrOnlyPositive = errors.New("invalid number")

func Classify(number int64) (result Classification, err error) {
	if number <= 0 {
		return result, ErrOnlyPositive
	}
	var i int64
	var sum int64
	for i = 1; i < number; i++ {
		if number%i == 0 {
			sum = sum + i
		}
	}
	if sum == number {
		result = ClassificationPerfect
	}
	if sum < number {
		result = ClassificationDeficient
	}
	if sum > number {
		result = ClassificationAbundant
	}
	return result, nil
}
