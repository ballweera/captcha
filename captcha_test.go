package captcha_test

import (
	"testing"

	"github.com/ballweera/captcha"
)

func TestPattern1(t *testing.T) {
	t.Run("1 + ONE", func(t *testing.T) {
		got := captcha.Generate(1, 1, 1, 1)
		want := "1 + ONE"
		assertError(want, got, t)
	})

	t.Run("1 + TWO", func(t *testing.T) {
		got := captcha.Generate(1, 1, 1, 2)
		want := "1 + TWO"
		assertError(want, got, t)
	})

	t.Run("1 + THREE", func(t *testing.T) {
		got := captcha.Generate(1, 1, 1, 3)
		want := "1 + THREE"
		assertError(want, got, t)
	})

	t.Run("1 + NINE", func(t *testing.T) {
		got := captcha.Generate(1, 1, 1, 9)
		want := "1 + NINE"
		assertError(want, got, t)
	})
}

func assertError(want, got string, t *testing.T) {
	if want != got {
		t.Errorf("want '%s' but got '%s'", want, got)
	}
}
