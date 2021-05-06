package summultiples

func SumMultiples(limit int, divisors ...int) int {
	sum := 0
	for i := 1; i < limit; i++ {
		for j := 0; j < len(divisors); j++ {
			if divisors[j] != 0 && i%divisors[j] == 0 {
				sum = sum + i
				break
			}
		}
	}
	return sum
}
