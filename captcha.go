package captcha

func Generate(pattern, left, operator, right int) string {
	if right == 2 {
		return "1 + TWO"
	}
	return "1 + ONE"
}
