package prime

func Nth(n int) (p int, s bool) {

	if n <= 0 {
		return 0, false
	}
	primecount := 0
	nth := 0
	s = true

	for i := 2; primecount < n; i++ {
		for j := 2; j < i; j++ {
			if i%j == 0 {
				s = false
				break
			} else {
				s = true
			}
		}
		if s {
			primecount++
			nth = i
		}
	}
	return nth, s
}
