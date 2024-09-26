// 🚀 Optical is a CLI tool that generates a Go Fiber project template. It is inspired by express-generator, a tool that generates a Node.js project template for the Express.js framework.

// 🚀 Optical generates a project structure with a basic setup for a Go Fiber API project, including a main.go file, handlers, middleware, models, routes, services, and a config package.

package main

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/charmbracelet/huh"
)

var (
	 version = "0.2.0"
	 logo = `
 _______         __   __              __ 
|       |.-----.|  |_|__|.----.---.-.|  |
|   -   ||  _  ||   _|  ||  __|  _  ||  |
|_______||   __||____|__||____|___._||__|
         |__|                            
     `
)
func main() {
	createFlag := flag.String("create","", "Creates a new Optical Project")
		versionFlag := flag.Bool("version", false, "Print the version of Optical CLI")
		flag.Bool("help", false, "Help for Optical CLI")
	
		flag.Usage = func() {
			fmt.Fprintf(os.Stderr, "Optical CLI - A tool for generating Go Fiber API projects\n\n")
			fmt.Fprintf(os.Stderr, "Usage:\n")
			fmt.Fprintf(os.Stderr, "optical create\n")
			fmt.Fprintf(os.Stderr, "optical [flags]\n\n")
			fmt.Fprintf(os.Stderr, "Flags:\n")
			flag.PrintDefaults()
		}
	
		flag.Parse()
	
		if *versionFlag {
			fmt.Printf("Optical CLI version %s\n", version)
			return
		}

	

	type Prompts struct{
		projectName string;
		ghUserName string;
		dbType  string;

	}

	prompts := Prompts{}

	if *createFlag != "" {
		println(*createFlag)
		fmt.Println(logo)

		form := huh.NewForm(
			huh.NewGroup(
				huh.NewInput().Title("What is the name of your project?").Value(&prompts.projectName).Validate(func(str string) error {
					if str == "" {
						return errors.New("project name can't be empty")

					}
					return nil
				}),
			),
		)

		err := form.Run()

		if err != nil {
			 fmt.Println("There is a error creating your project")
			 return
		}

		fmt.Println(prompts.projectName)


		
	} else {
		fmt.Println(logo)
		flag.Usage()

		
		os.Exit(1)
	}
	
}

// func getCurrentDirectory() string {
// 	dir, err := os.Getwd()
// 	if err != nil {
// 		fmt.Println("❗Error getting current directory:", err)
// 		os.Exit(1)
// 	}
// 	return dir
// }

