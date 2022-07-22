package config

import (
	"log"
	"text/template"
)

type AppConfig struct {
	InfoLogger      *log.Logger
	ErrorLogger     *log.Logger
	AppInProduction bool
	UseTempCache    bool
	// Session         *scs.SessionManager
	TemplateCache   map[string]*template.Template
}
