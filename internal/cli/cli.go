package cli

import "fmt"

const HELP string = "-h"

func ParseArgs(args []string) string {
	option := args[1]
	switch option {
	case HELP:
		return "HELP"
	default:
		return fmt.Sprintf("%s is not a valid option", option)
	}
}