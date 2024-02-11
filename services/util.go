package services

// Adjusting numbers into same lengths by adding leading zeros
func numToSameLen(a, b string) (string, string) {
	var excDigits int
	if len(a) >= len(b) {
		excDigits = len(a) - len(b)
		for i := 0; i < excDigits; i++ {
			b = `0` + b
		}
	} else {
		excDigits = len(b) - len(a)
		for i := 0; i < excDigits; i++ {
			a = `0` + a
		}
	}

	return a, b
}
