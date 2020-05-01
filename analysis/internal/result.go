package internal

import "strings"

type Result struct {
	Imports   []string
	List      string
	Package   string
	Type      string
	Map       MapData
	CountType string
}

func NewResult(countType string) *Result {
	return &Result{
		CountType: countType,
	}
}

func (r *Result) SetList(list []string) {
	r.List = strings.Join(list, ", ")
}

func (r *Result) SetPackage(_package string) {
	r.Package = _package
}

func (r *Result) SetMapData(md MapData) {
	r.Map = md
}

func (r *Result) SetImports(imps []string) {
	r.Imports = imps
}
