package cli

import (
	"errors"
	"fmt"
	"strconv"
)

var ( 
	port string = "-p"

	ErrMissingPort error = errors.New("Missing port number expected")
)

func ParsePort(args []string) (int, error) {
	if len(args) < 3 {
		return -1, ErrMissingPort
	}
	portNum, err := strconv.ParseInt(args[2], 10, 64)
	if err != nil {
		return -1, err
	}
	return int(portNum), nil
}

func ParseArgs(args []string) (int, error) {
	option := args[1]
	switch option {
	case port:
		return ParsePort(args)
	default:
		return -1, fmt.Errorf("%s is not a valid option", option)
	}
}
