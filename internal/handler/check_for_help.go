package handler

import (
	"fmt"
	"please_generator/internal/info"
)

func CheckForHelp(args []string) {
	helpCommands := map[string]bool{"help": true, "h": true}
	argsLen := len(args)
	var ok bool

	if argsLen != 0 {
		_, ok = helpCommands[args[0]]
	}
	if argsLen == 0 || ok == true {
		fmt.Println(string(info.GLOBALINFO))
	}
}
