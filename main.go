// 🚀 Optical is a CLI tool that generates a Go Fiber project template. It is inspired by express-generator, a tool that generates a Node.js project template for the Express.js framework.

// 🚀 Optical generates a project structure with a basic setup for a Go Fiber API project, including a main.go file, handlers, middleware, models, routes, services, and a config package.

package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/TuhinBar/optical/lib/generator"
	"github.com/TuhinBar/optical/lib/helper"
)

var (
	 version = "0.5.0"
	 logo = `
 _______         __   __              __ 
|       |.-----.|  |_|__|.----.---.-.|  |
|   -   ||  _  ||   _|  ||  __|  _  ||  |
|_______||   __||____|__||____|___._||__|
         |__|                            `
)




func main() {
	createFlag := flag.Bool("create",false, "Creates a new Optical Project")
	versionFlag := flag.Bool("version", false, "Print the version of Optical CLI")
	flag.Bool("help", false, "Help for Optical CLI")

	flag.Usage = func() {
		helper.PrintFlags()
	}

	flag.Parse()
	
	if *versionFlag {
		helper.PrintVersion(version)
		return
	}



	if *createFlag {
		fmt.Println(logo)
		helper.PrintVersion(version)

		form,prompts := helper.CreateForm()
		
		err := form.Run()

		if err != nil {
			 fmt.Println("There is an error creating your project")
			 return
		}

		if prompts.ProjectName == "./" || prompts.ProjectName == "" {
			prompts.ProjectName = filepath.Base(helper.GetCurrentDirectory())
			
		}

		projectErr := generator.GenerateProject(prompts.ProjectName,".",prompts.GhUserName)

		if projectErr != nil {
			fmt.Printf("❗Error generating project: %v\n", projectErr)
			os.Exit(1)
		}

		helper.PrintSuccessMsg(prompts.ProjectName)

	} else {
		fmt.Println(logo)
		flag.Usage()

		
		os.Exit(1)
	}
	
}

