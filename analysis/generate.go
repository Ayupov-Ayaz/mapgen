package analysis

import (
	"fmt"
	"text/template"

	"github.com/ayupov-ayaz/mapgen/templates"
)

type Writer interface {
	WriteToFile(filename string, t *template.Template, i interface{}) error
}

func prepareTemplate(templatePath string) *template.Template {
	return template.Must(template.New("const-list").Parse(templatePath))
}

func GenerateMapByString(w Writer, data MapParams) error {
	result, err := analysisFileByMap(data)
	if err != nil {
		return err
	}

	result.SetPackage(data.PackageName)
	result.Type = data.StructName
	t := prepareTemplate(templates.MapImpl)

	if err := w.WriteToFile(data.FilePath, t, result); err != nil {
		return fmt.Errorf("write to file failed: %w", err)
	}

	return nil
}
