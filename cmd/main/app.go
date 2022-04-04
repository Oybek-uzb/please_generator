package main

import (
	"os"
	"please_generator/internal/handler"
)

func main() {
	args := os.Args[1:]

	handler.CheckForHelp(args)

}
