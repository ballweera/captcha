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
}
