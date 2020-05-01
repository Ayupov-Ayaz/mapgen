package analysis

import "strings"

type MapParams struct {
	FilePath  string
	CountType string
}

func NewMapParams(filepath, countType string) MapParams {
	return MapParams{
		CountType: countType,
		FilePath:  strings.Replace(filepath, ".go", "", 1) + ".go",
	}
}
