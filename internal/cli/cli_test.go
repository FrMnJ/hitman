package cli

import "testing"

func TestParseArgs(t *testing.T) {
	t.Run("invalid option must return error message", func(t *testing.T) {
		given := []string{"hitman", "-d"}
		got := ParseArgs(given)
		expected := "-d is not a valid option"

		if got != expected {
			t.Errorf("expected %q but got %q, when given %v", expected, got, given)
		}
	})
}