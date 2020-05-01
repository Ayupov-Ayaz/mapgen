package analysis

import "strings"

type MapParams struct {
	PackageName string
	FilePath    string
	MapType     string
	CountType   string
}

func NewMapParams(packageName, filepath, mapType, countType string) MapParams {
	return MapParams{
		PackageName: packageName,
		MapType:     mapType,
		CountType:   countType,
		FilePath:    strings.Replace(filepath, ".go", "", 1) + ".go",
	}
}
