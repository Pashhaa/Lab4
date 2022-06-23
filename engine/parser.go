package engine

import (
	"strconv"
	"strings"
)

func Parse(toParse string) command {
	args := strings.Fields(toParse)

	if len(args) < 2 {
		return &printCommand{"syntax error: not enough arguments"}
	}

	cmd := args[0]
	cmdargs := args[1:]

	switch cmd {
	case "print":
		message := strings.Join(cmdargs, " ")
		return &printCommand{message}

	case "printc":
		if len(cmdargs) != 2 {
			return &printCommand{"syntax error: invalid arguments"}
		}
		tempvar, _ := strconv.Atoi(cmdargs[1])
		return &PrintcCommand{cmdargs[0], tempvar}

	default:
		return &printCommand{"syntax error: invalid command"}
	}
}
