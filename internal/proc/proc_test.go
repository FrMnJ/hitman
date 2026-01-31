package proc

import (
	"errors"
	"testing"
)

func TestSpawnProcOnPort(t *testing.T) {
	t.Run("given a invalid port returns error", func(t *testing.T) {
		given := 77777
		_, err := SpawnProcOnPort(given)
		if !errors.Is(err, ErrInvalidPort) {
			t.Errorf("expected %q but got %q, given %d", ErrInvalidPort, err, given)
		}
	})
}