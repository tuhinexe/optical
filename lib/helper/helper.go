package helper

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/charmbracelet/huh"
)
type Prompts struct{
	ProjectName string;
	GhUserName string;
	HasAir bool
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
			huh.NewInput().Title("What is the name of your project?(e.g 'my-fiber-api','./')").Description("A new folder with this name will be created.").Value(&prompts.ProjectName).Validate(func(str string) error {
				if str == "" {
					return errors.New("project name can't be empty")

				}
				return nil
			}),
			huh.NewInput().Title("Enter your GitHub username.").Description("This is required to create the go.mod file").Value(&prompts.GhUserName).Validate(func(s string) error {
				if s == "" {
					return errors.New("github username is required")
				}
				return nil
			}),
			huh.NewSelect[bool]().Title("Do you have air installed ?").Description("Air is required for auto-reload").Options(
				huh.NewOption("Yes",true),
				huh.NewOption("No",false),
			).Value(&prompts.HasAir),
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
3. air`)
}

func ShowLoadingIndicator(done chan bool) {
	loadingChars := []rune{'|', '/', '-', '\\'}
	i := 0
	for {
		select {
		case <-done:
			fmt.Print("\r")
			return
		default:
			fmt.Printf("\rCreating your project... %c", loadingChars[i%len(loadingChars)])
			i++
			time.Sleep(100 * time.Millisecond)
		}
	}
}


func FetchTemplateFromGithub(templateName string) (string,error) {
	fileUrl := fmt.Sprintf("https://raw.githubusercontent.com/TuhinBar/optical/main/lib/templates/%s", templateName)

	resp, err := http.Get(fileUrl)

	if err != nil {
		return "", fmt.Errorf("❗Failed to download template %s: %w", templateName, err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("❗Failed to download template %s: HTTP status %d", templateName, resp.StatusCode)
	}

	body,err := io.ReadAll(resp.Body)

	if err != nil {
		return "",fmt.Errorf("❗Failed to read response body for template %s: %w", templateName, err)
	}

	return string(body),nil
}