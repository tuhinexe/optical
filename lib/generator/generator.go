package generator

import (
	"embed"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"text/template"

	"github.com/TuhinBar/optical/lib/helper"
)

//go:embed templates/*.tmpl
var templateFS embed.FS


func GenerateProject(name, path string,ghUsername string,hasAir bool) error {
	projectPath := filepath.Join(path, name)

	if err := os.MkdirAll(projectPath, os.ModePerm); err != nil {
		return fmt.Errorf("❗ Failed to create project directory: %w", err)
	}

	dirs := []string{ "handlers", "middleware", "models","routes","services", "config"}
	for _, dir := range dirs {
		if err := os.MkdirAll(filepath.Join(projectPath, dir), os.ModePerm); err != nil {
			return fmt.Errorf("❗Failed to create directory %s: %w", dir, err)
		}
	}

	files := map[string]string{
		"main.go":                  "main.go.tmpl",
		"handlers/handler.go":   	"handler.go.tmpl",
		"middleware/middleware.go": "middleware.go.tmpl",
		"services/service.go":   	"service.go.tmpl",
		"models/models.go":      	"models.go.tmpl",
		"routes/routes.go":      	"routes.go.tmpl",
		"config/config.go":      	"config.go.tmpl",
		"go.mod":                	"go.mod.tmpl",
		".air.toml":			 	".air.toml.tmpl",
	}

	for filePath, templateName := range files {
		if err := generateFileFromTemplate(projectPath, filePath, templateName, name,ghUsername); err != nil {
			return fmt.Errorf("❗Failed to create %s: %w", filePath, err)
		}
	};
	
	if !hasAir {
		done := make(chan bool)
		go helper.ShowLoadingIndicator(done)
		cmd := exec.Command("go", "install", "github.com/air-verse/air@latest")
		cmd.Dir = projectPath
	
		if output, err := cmd.CombinedOutput(); err != nil {
			done <- true
			fmt.Printf("❗Failed to install air in your project: %v\nOutput: %s\n", err, string(output))
		} else {
			done <- true
			fmt.Println("✅ Successfully installed air for auto-reload")
		}
	}

	return nil
}

func generateFileFromTemplate(projectPath, filePath, templateName, projectName string,ghUsername string) error {
	fullPath := filepath.Join(projectPath, filePath)

	tmplContent, err := templateFS.ReadFile(filepath.Join("templates", templateName))
	if err != nil {
		return fmt.Errorf("❗ Failed to read template %s: %w", templateName, err)
	}

	tmpl, err := template.New(templateName).Parse(string(tmplContent))
	if err != nil {
		return fmt.Errorf("❗Failed to parse template %s: %w", templateName, err)
	}

	file, err := os.Create(fullPath)
	if err != nil {
		return fmt.Errorf("❗Failed to create file %s: %w", filePath, err)
	}
	defer file.Close()

	data := struct {
		ProjectName string
		GitHubUsername string
	}{
		ProjectName: projectName,
		GitHubUsername: ghUsername,
	}

	if err := tmpl.Execute(file, data); err != nil {
		return fmt.Errorf("❗Failed to execute template %s: %w", templateName, err)
	}

	return nil
}