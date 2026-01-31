package proc

import "errors"

const (
	MaxPort int = 65535
	MinPort int = 0
	InvalidPID int = -1
)

var (
	ErrInvalidPort error = errors.New("Invalid port range 0 to 65535")
	NotImplemented error = errors.New("Not Implemented")
)

func SpawnProcOnPort(port int) (pid int, err error) {
	if port > MaxPort || port < MinPort {
		return InvalidPID, ErrInvalidPort
	}
	return InvalidPID, NotImplemented
}