package templates

const MapImpl = `//CODE GENERATED AUTOMATICALLY. DO NOT EDIT.
package {{.Package}}

import (
{{range $k, $val := $.Imports}}
{{$val}}{{end}}	
)

type {{$.Type}} struct {}

func (p {{$.Type}}) Get(s {{$.Map.KV.Key.FullType}}, count {{$.CountType}}) {{$.Map.KV.Val.FullType}} {
	switch s {
		{{range $key, $val := $.Map.Data }}
			case {{$key}}:
				if count < {{$val.Length}} && count >= 0  {
					return {{"[]"}}{{$.Map.KV.Val.FullType}}{{"{"}}{{$val.Join}}{{"}"}}[count]
				}
		{{end}}
	}
	return 0
}
`
