package cli

import (
	"fmt"
	"strconv"
)

const PORT string = "-p"

func ParseArgs(args []string) (int, error) {
	option := args[1]
	switch option {
	case PORT:
		port := args[2]
		portNum, err := strconv.ParseInt(port, 10, 64)
		if err != nil{
			return -1, err
		}
		return int(portNum), nil
	default:
		return -1, fmt.Errorf("%s is not a valid option", option)
	}
}