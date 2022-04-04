package handler

import "fmt"

func CheckForCommand(args []string) {
	if len(args) != 0 {
		switch args[0] {
		case "generate", "gen":
			fmt.Println("ok, generating...")
		case "help", "h":
			fmt.Println("ok, helping...")
		case "add":
			fmt.Println("ok, adding...")
		}
	}
}
