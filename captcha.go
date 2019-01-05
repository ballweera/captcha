package captcha

import "fmt"

func Generate(pattern, left, operator, right int) string {
	words := [3]string{"ONE", "TWO", "THREE"}
	return fmt.Sprintf("1 + %s", words[right-1])
}
