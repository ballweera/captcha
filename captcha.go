package captcha

import "fmt"

func Generate(pattern, left, operator, right int) string {
	words := []string{"ONE", "TWO", "THREE", "FOUR", "FIVE", "SIX", "SEVEN", "EIGHT", "NINE"}
	var result string

	switch pattern {
	case 1:
		result = fmt.Sprintf("1 + %s", words[right-1])
	case 2:
		result = fmt.Sprintf("%s + 1", words[left-1])
	}
	return result
}
