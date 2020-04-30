package internal

import "strings"

type Result struct {
	DataList []string
	List     string
	Package  string
	Type     string
	Map      MapData
}

func NewResult(Type string) *Result {
	return &Result{
		Type: Type,
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
