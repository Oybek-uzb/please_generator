package main

import (
	"os"
	"please_generator/internal/handler"

	"github.com/manifoldco/promptui"
	"github.com/sirupsen/logrus"
)

func main() {
	args := os.Args[1:]

	handler.CheckForHelp(args)

}

func stringPrompt(label, defaultValue string) string {
	var (
		err    error
		result string
	)

	prompt := promptui.Prompt{
		Label:   label,
		Default: defaultValue,
	}

	if result, err = prompt.Run(); err != nil {
		logrus.WithError(err).Fatal("error asking for '%s'", label)
	}

	return result
}
