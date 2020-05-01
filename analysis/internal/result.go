package internal

import "strings"

type Results struct {
	Package string
	Imports []string
	R       []Result
}

func getUniqueImports(r []Result) []string {
	var imports []string
	if len(r) > 0 {
		imports = r[0].Imports
	}

	for i := 1; i < len(r); i++ {
		for _, imp := range r[i].Imports {
			ok := true
			for _, reqImp := range imports {
				if imp == reqImp {
					ok = false
					break
				}
			}

			if ok {
				imports = append(imports, imp)
			}
		}
	}

	return imports
}

func NewResults(_package string, results []Result) *Results {
	return &Results{
		Package: _package,
		Imports: getUniqueImports(results),
		R:       results,
	}
}

type Result struct {
	Imports    []string
	List       string
	StructName string
	KeyType    string
	ValType    string
	Map        MapData
	CountType  string
	Condition  string
}

func NewResult(countType string) *Result {
	condition := ""
	switch countType {
	case "int8", "int16", "int32", "int64":
		condition = "&& count >= 0"
	}

	return &Result{
		Condition: condition,
		CountType: countType,
	}
}

func (r *Result) SetList(list []string) {
	r.List = strings.Join(list, ", ")
}

func (r *Result) SetMapData(md MapData) {
	r.KeyType = md.KV.Key.FullType
	r.ValType = md.KV.Val.FullType

	r.Map = md
}

func (r *Result) SetImports(imps []string) {
	r.Imports = imps
}
