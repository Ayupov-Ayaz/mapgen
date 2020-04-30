package analysis

import (
	"fmt"
	"text/template"

	"github.com/ayupov-ayaz/mapgen/templates"
)

type TemplateRecorder interface {
	RecordToFile(filename string, t *template.Template, i interface{}) error
}

func prepareTemplate(templatePath string) *template.Template {
	return template.Must(template.New("const-list").Parse(templatePath))
}

func GenerateMapByString(w TemplateRecorder, data MapParams) error {
	result, err := analysisFileByMap(data)
	if err != nil {
		return err
	}

	result.SetPackage(data.PackageName)
	result.Type = data.StructName
	t := prepareTemplate(templates.MapImpl)

	if err := w.RecordToFile(data.FilePath, t, result); err != nil {
		return fmt.Errorf("write to file failed: %w", err)
	}

	return nil
}
