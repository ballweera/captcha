package captcha

import (
	"fmt"
	"strconv"
)

var words []string = []string{"ZERO", "ONE", "TWO", "THREE", "FOUR", "FIVE", "SIX", "SEVEN", "EIGHT", "NINE"}
var operatorWords []string = []string{"+", "-", "*"}

func Generate(pattern, l, op, r int) string {
	return fmt.Sprintf("%s %s %s", leftOperand(pattern, l), operator(op), rightOperand(pattern, r))
}

func leftOperand(pattern, value int) string {
	if pattern == 1 {
		return strconv.Itoa(value)
	} else {
		return words[value]
	}
}

func rightOperand(pattern, value int) string {
	if pattern == 1 {
		return words[value]
	} else {
		return strconv.Itoa(value)
	}
}

func operator(value int) string {
	return operatorWords[value-1]
}
