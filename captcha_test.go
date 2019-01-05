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

	t.Run("5 + FOUR", func(t *testing.T) {
		got := captcha.Generate(1, 5, 1, 4)
		want := "5 + FOUR"
		assertError(want, got, t)
	})

	t.Run("8 - TWO", func(t *testing.T) {
		got := captcha.Generate(1, 8, 2, 2)
		want := "8 - TWO"
		assertError(want, got, t)
	})

	t.Run("4 * FOUR", func(t *testing.T) {
		got := captcha.Generate(1, 4, 3, 4)
		want := "4 * FOUR"
		assertError(want, got, t)
	})
}

func TestPattern2(t *testing.T) {
	t.Run("ONE + 1", func(t *testing.T) {
		got := captcha.Generate(2, 1, 1, 1)
		want := "ONE + 1"
		assertError(want, got, t)
	})

	t.Run("NINE + 1", func(t *testing.T) {
		got := captcha.Generate(2, 9, 1, 1)
		want := "NINE + 1"
		assertError(want, got, t)
	})

	t.Run("TWO + 5", func(t *testing.T) {
		got := captcha.Generate(2, 2, 1, 5)
		want := "TWO + 5"
		assertError(want, got, t)
	})
}

func assertError(want, got string, t *testing.T) {
	if want != got {
		t.Errorf("want '%s' but got '%s'", want, got)
	}
}
