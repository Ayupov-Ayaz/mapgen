package templates

const MapImpl = `//CODE GENERATED AUTOMATICALLY. DO NOT EDIT.
package {{.Package}}

import ({{range $k, $val := $.Imports}}
{{$val}}{{end}}	
)
{{range $i, $result := $.R}}
func {{$result.FuncName}} (s {{$result.KeyType}}, count {{$result.CountType}}) {{$result.ValType}} {
	switch s {
		{{range $key, $val := $result.Map.Data }}
			case {{$key}}:
				if count < {{$val.Length}}{{$result.Condition}}{
					return {{"[]"}}{{$result.ValType}}{{"{"}}{{$val.Join}}{{"}"}}[count]}
		{{end}}
	}
	return 0
}

{{end}}
`
