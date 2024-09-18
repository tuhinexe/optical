package generator

import (
	"embed"
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

//go:embed templates/*.tmpl
var templateFS embed.FS


func GenerateProject(name, path string,ghUsername string) error {
	projectPath := filepath.Join(path, name)

	if err := os.MkdirAll(projectPath, os.ModePerm); err != nil {
		return fmt.Errorf("❗ Failed to create project directory: %w", err)
	}

	dirs := []string{"cmd", "internal/handlers", "internal/middleware", "internal/models","internal/routes","internal/services", "config"}
	for _, dir := range dirs {
		if err := os.MkdirAll(filepath.Join(projectPath, dir), os.ModePerm); err != nil {
			return fmt.Errorf("❗Failed to create directory %s: %w", dir, err)
		}
	}

	files := map[string]string{
		"cmd/main.go":                    "main.go.tmpl",
		"internal/handlers/handler.go":   "handler.go.tmpl",
		"internal/middleware/middleware.go": "middleware.go.tmpl",
		"internal/services/service.go":   "service.go.tmpl",
		"internal/models/models.go":      "models.go.tmpl",
		"internal/routes/routes.go":      "routes.go.tmpl",
		"config/config.go":               "config.go.tmpl",
		"go.mod":                         "go.mod.tmpl",
	}

	for filePath, templateName := range files {
		if err := generateFileFromTemplate(projectPath, filePath, templateName, name,ghUsername); err != nil {
			return fmt.Errorf("❗Failed to create %s: %w", filePath, err)
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