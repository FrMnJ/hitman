package proc

import "errors"

var (
	ErrInvalidPort error = errors.New("Invalid port range 0 to 65535")
	NotImplemented error = errors.New("Not Implemented")
)

func SpawnProcOnPort(port int) (pid int, err error) {
	if port > 65535 || port < 0 {
		return -1, ErrInvalidPort
	}
	return -1, NotImplemented
}