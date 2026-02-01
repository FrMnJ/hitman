package proc

import (
	"errors"
	"testing"
)

func TestKillProcOnPort(t *testing.T) {
	t.Run("given a invalid port returns error", func(t *testing.T) {
		given := 77777
		err := KillProcOnPort(given)
		if !errors.Is(err, ErrInvalidPort) {
			t.Errorf("expected %q but got %q, given %d", ErrInvalidPort, err, given)
		}
	})
}