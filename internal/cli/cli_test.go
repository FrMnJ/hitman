package cli

import (
	"errors"
	"testing"
)

func TestParseArgs(t *testing.T) {
	t.Run("invalid option must return error message", func(t *testing.T) {
		given := []string{"hitman", "-d"}
		_, err := ParseArgs(given)

		if !errors.Is(err, ErrInvalidOption) {
			t.Errorf("expected %q but got %q, when given %v", ErrInvalidOption, err, given)
		}
	})

	t.Run("given a port parameter returns the port", func(t *testing.T) {
		given := []string{"hitman", "-p", "8080"}
		expected := 8080
		got, err := ParseArgs(given)

		if err != nil {
			t.Errorf("error must be nil but got %q", err)
		}

		if got != expected {
			t.Errorf("expected %d but got %d, when given %v", expected, got, given)
		}
	})

	t.Run("not giving port parameter returns error", func(t *testing.T) {
		given := []string{"hitman", "-p"}
		_, err := ParseArgs(given)
		if !errors.Is(err, ErrMissingPort){
			t.Errorf("expected %q but got %q", ErrMissingPort, err)
		}
	})
}