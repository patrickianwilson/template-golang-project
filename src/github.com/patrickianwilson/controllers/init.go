package controllers

import (
	"github.com/InquestDevops/templater"
	"os"
)

var templateManager templater.TemplateManager

func init() {
	templatePath := os.Getenv("TEMPLATE_PATH")
	templateManager = templater.NewBuiltinTemplateManagerWithDefaultIncludes(templatePath, []string{"common/layout"})
}
