package analysis

import "strings"

type MapParams struct {
	PackageName string
	FilePath    string
	MapType     string
	StructName  string
}

func NewMapParams(packageName, filepath, mapType, structName string) MapParams {
	return MapParams{
		PackageName: packageName,
		MapType:     mapType,
		StructName:  structName,
		FilePath:    strings.Replace(filepath, ".go", "", 1) + ".go",
	}
}
