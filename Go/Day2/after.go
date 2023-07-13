package Day2

func FilterOddNumbers(numbers []int) []int {
	var result []int
	for _, number := range numbers {
		if number%2 == 0 {
			result = append(result, number)
		}
	}
	return result
}

type Number interface {
	int | float64 | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32
}

func SquareNumbers[T Number](numbers []T) []T {
	var result []T
	for _, number := range numbers {
		result = append(result, number*number)
	}
	return result
}

type Measurable interface {
	string | []any
}

func CountChars[T Measurable](s []T) []int {
	var result []int
	for _, str := range s {
		result = append(result, len(str))
	}
	return result
}
