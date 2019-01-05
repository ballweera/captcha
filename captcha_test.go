package captcha_test

import (
	"testing"

	"github.com/ballweera/captcha"
)

func TestPattern1(t *testing.T) {
	t.Run("1 + ONE", func(t *testing.T) {
		got := captcha.Generate(1, 1, 1, 1)
		want := "1 + ONE"

		if want != got {
			t.Errorf("want '%s' but got '%s'", want, got)
		}
	})

	t.Run("1 + TWO", func(t *testing.T) {
		got := captcha.Generate(1, 1, 1, 2)
		want := "1 + TWO"

		if want != got {
			t.Errorf("want '%s' but got '%s'", want, got)
		}
	})

	t.Run("1 + THREE", func(t *testing.T) {
		got := captcha.Generate(1, 1, 1, 3)
		want := "1 + THREE"

		if want != got {
			t.Errorf("want '%s' but got '%s'", want, got)
		}
	})

	t.Run("1 + NINE", func(t *testing.T) {
		got := captcha.Generate(1, 1, 1, 9)
		want := "1 + NINE"

		if want != got {
			t.Errorf("want '%s' but got '%s'", want, got)
		}
	})
}
