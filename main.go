// 🚀 Optical is a CLI tool that generates a Go Fiber project template. It is inspired by express-generator, a tool that generates a Node.js project template for the Express.js framework.

// 🚀 Optical generates a project structure with a basic setup for a Go Fiber API project, including a main.go file, handlers, middleware, models, routes, services, and a config package.

package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/TuhinBar/optical/lib/generator"
)



var (
	 version = "0.1.0"
	 logo = `
 _______         __   __              __ 
|       |.-----.|  |_|__|.----.---.-.|  |
|   -   ||  _  ||   _|  ||  __|  _  ||  |
|_______||   __||____|__||____|___._||__|
         |__|                            
     `
)
func main() {
		initFlag := flag.Bool("init", false, "Initialize a new Optical project in the current directory")
		projectName := flag.String("name", "", "Name of the project")
		versionFlag := flag.Bool("version", false, "Print the version of Optical CLI")
	
		flag.Usage = func() {
			fmt.Fprintf(os.Stderr, "Optical CLI - A tool for generating Go Fiber API projects\n\n")
			fmt.Fprintf(os.Stderr, "Usage:\n")
			fmt.Fprintf(os.Stderr, "optical -init -name '<your-project-name>'\n")
			fmt.Fprintf(os.Stderr, "optical [flags]\n\n")
			fmt.Fprintf(os.Stderr, "Flags:\n")
			flag.PrintDefaults()
		}
	
		flag.Parse()
	
		// Check for version flag
		if *versionFlag {
			fmt.Printf("Optical CLI version %s\n", version)
			return
		}

	if *initFlag {
		name := *projectName
		fmt.Println(logo)
		if name == "" {
			name = filepath.Base(getCurrentDirectory())
		}
		err := generator.GenerateProject(name, ".")
		if err != nil {
			fmt.Printf("❗Error generating project: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("🎉 Optical project created successfully\n")

		fmt.Println(`
To run the project follow these steps:
1. cd ` + name + `
2. go mod tidy
3. go run ./cmd/main.go
		`)
	} else {

		
	fmt.Println(logo)

		flag.Usage()
		os.Exit(1)

		

		
		

	  
	}
}

func getCurrentDirectory() string {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return dir
}

