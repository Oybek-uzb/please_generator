package main

import (
	"embed"
	"fmt"
	"os"
	"text/template"

	"github.com/manifoldco/promptui"
	"github.com/sirupsen/logrus"
)

var (
	// go:embed templates/*.tmpl
	rootFs embed.FS
)

type appValue struct {
	AppName  string
	YourName string
}

func main() {
	var (
		err       error
		fp        *os.File
		templates *template.Template
		subdirs   []string
	)
	values := appValue{}
	fmt.Printf("Welcome to the Please Generator App")

	values.AppName = stringPrompt("Enter this application's name (no spaces)", "")
	values.YourName = stringPrompt("Enter Your name", "")

	rootFsMapping := map[string]string{
		"index.html.tmpl": "static/index.html",
		"main.go.tmpl":    "main.go",
	}

	// Creating directories
	if err = os.Mkdir(values.AppName, 0755); err != nil {
		logrus.WithError(err).Errorf("error attempting to create application directory '%s'", values.AppName)
	}

	if err = os.Chdir(values.AppName); err != nil {
		logrus.WithError(err).Errorf("error changing to new directory '%s'", values.AppName)
	}

	subdirs = []string{
		"static",
	}

	for _, dirname := range subdirs {
		if err = os.MkdirAll(dirname, 0755); err != nil {
			logrus.WithError(err).Errorf("unable to create subdirectory '%s'", dirname)
		}
	}

	// Process templates
	if templates, err = template.ParseFiles(rootFs, "templates/*.tmpl"); err != nil {
		logrus.WithError(err).Fatal("error parsing root templates files")
	}

	for templateName, outputPath := range rootFsMapping {
		if fp, err = os.Create(outputPath); err != nil {
			logrus.WithError(err).Fatalf("unable to create file %s for writing", outputPath)
		}

		defer fp.Close()

		if err = templates.ExecuteTemplate(fp, templateName, values); err != nil {
			logrus.WithError(err).Fatalf("unable to exeucte template %s", templateName)
		}
	}

	fmt.Printf("\n🎉 Congratulations! Your new application is ready.")
  	fmt.Printf("\nTo begin execute the following:\n\n")
	fmt.Printf("   cd %s\n", values.AppName)
  	fmt.Printf("   go run .\n")
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
