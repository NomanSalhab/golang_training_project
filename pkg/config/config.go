package config

import (
	"html/template"
	"log"
)

// AppConfig holds The Application Config
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	infoLog       *log.Logger
}
