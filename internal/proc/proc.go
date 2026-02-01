package proc

import "errors"

const (
	MaxPort int = 65535
	MinPort int = 0
)

var (
	ErrInvalidPort error = errors.New("Invalid port range 0 to 65535")
	NotImplemented error = errors.New("Not Implemented")
)

func KillProcOnPort(port int) error {
	if port > MaxPort || port < MinPort {
		return ErrInvalidPort
	}
	return NotImplemented
}