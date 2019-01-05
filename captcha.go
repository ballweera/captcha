package captcha

import "fmt"

func Generate(pattern, left, operator, right int) string {
	words := []string{"ONE", "TWO", "THREE", "FOUR", "FIVE", "SIX", "SEVEN", "EIGHT", "NINE"}
	return fmt.Sprintf("1 + %s", words[right-1])
}
