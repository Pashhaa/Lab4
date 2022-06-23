package engine

import (
	"strconv"
	"strings"
)

func Parse(toParse string) Command {
	args := strings.Fields(toParse)

	if len(args) < 2 {
		return &PrintCommand{"syntax error: not enough arguments"}
	}

	cmd := args[0]
	cmdargs := args[1:]

	switch cmd {
	case "print":
		message := strings.Join(cmdargs, " ")
		return &PrintCommand{message}

	case "printc":
		if len(cmdargs) != 2 {
			return &PrintCommand{"syntax error: invalid arguments"}
		}
		tempvar, _ := strconv.Atoi(cmdargs[1])
		return &PrintcCommand{cmdargs[0], tempvar}

	default:
		return &PrintCommand{"syntax error: invalid command"}
	}
}
