package analysis

import "strings"

type MapParams struct {
	PackageName string
	FilePath    string
	MapType     string
	StructName  string
	CountType   string
}

func NewMapParams(packageName, filepath, mapType, structName, countType string) MapParams {
	return MapParams{
		PackageName: packageName,
		MapType:     mapType,
		StructName:  structName,
		CountType:   countType,
		FilePath:    strings.Replace(filepath, ".go", "", 1) + ".go",
	}
}
