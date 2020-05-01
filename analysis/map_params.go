package analysis

import "strings"

type MapParams struct {
	FilePath  string
	MapType   string
	CountType string
}

func NewMapParams(filepath, mapType, countType string) MapParams {
	return MapParams{
		MapType:   mapType,
		CountType: countType,
		FilePath:  strings.Replace(filepath, ".go", "", 1) + ".go",
	}
}
