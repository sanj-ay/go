package secret

func Handshake(number uint) []string {
	handshakes := []string{"wink", "double blink", "close your eyes", "jump"}
	var result []string
	num := number
	if num >= 1<<4 {
		num = num - 1<<4

	}
	for num > 0 {
		for i := 0; i <= 3; i++ {

			if num > 1<<4 {
				num = num - 1<<4
			}
			if num == 1<<i {
				result = append(result, handshakes[i])
				num = 0
				break
			}

			if 1<<i < num && !(num >= 1<<(i+1)) {
				result = append(result, handshakes[i])
				num = num - 1<<(i)
				continue

			}

		}

	}
	if number < 16 {
		for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
			result[i], result[j] = result[j], result[i]
		}

	}
	return result

}
