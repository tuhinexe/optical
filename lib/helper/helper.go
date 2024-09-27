package helper

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/charmbracelet/huh"
)
type Prompts struct{
	ProjectName string;
	GhUserName string;
	DbType  string;

}


func PrintVersion(version string) {
	fmt.Printf("Optical CLI version %s\n", version)

}

func PrintFlags() {
	fmt.Fprintf(os.Stderr, "Optical CLI - A tool for generating Go Fiber API projects\n\n")
	fmt.Fprintf(os.Stderr, "Usage:\n")
	fmt.Fprintf(os.Stderr, "optical create\n")
	fmt.Fprintf(os.Stderr, "optical [flags]\n\n")
	fmt.Fprintf(os.Stderr, "Flags:\n")
	flag.PrintDefaults()
}


func GetCurrentDirectory() string {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("❗Error getting current directory:", err)
		os.Exit(1)
	}
	return dir
}

func CreateForm() (*huh.Form,*Prompts) {
	prompts := Prompts{}
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().Title("What is the name of your project?(e.g 'my-fiber-api','./')").Value(&prompts.ProjectName).Validate(func(str string) error {
				if str == "" {
					return errors.New("project name can't be empty")

				}
				return nil
			}),
			huh.NewInput().Title("Enter your GitHub username.").Value(&prompts.GhUserName).Validate(func(s string) error {
				if s == "" {
					return errors.New("github username is required")
				}
				return nil
			}),
		),
	)

	return form,&prompts
}


func PrintSuccessMsg(name string) {
	fmt.Printf("🎉 Optical project created successfully\n")

		fmt.Println(`
To run the project follow these steps:
1. cd ` + name + `
2. go mod tidy
3. go run ./cmd/main.go
		`)
}