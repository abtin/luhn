// Package luhn provides luhn validation. To learn about luhn algorithm see: https://en.wikipedia.org/wiki/Luhn_algorithm
package luhn

// Valid returns whether the input string is a valid luhn string
func Valid(input string) bool {
	var sum, counter int
	for i := len(input) - 1; i >= 0; i-- {
		chr := input[i]
		if chr == ' ' {
			continue
		}
		if chr < '0' || chr > '9' {
			return false
		}
		digit := int(chr - '0')
		if counter%2 != 0 {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}
		sum += digit
		counter++
	}
	if counter <= 1 {
		return false
	}
	return sum%10 == 0
}
