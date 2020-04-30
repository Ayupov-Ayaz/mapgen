package templates

const MapImpl = `//CODE GENERATED AUTOMATICALLY. DO NOT EDIT.
package {{.Package}}

type {{$.Type}} struct {}

func (p {{$.Type}}) Get(i {{$.Map.KeyType}}, count int) {{$.Map.ValType}} {
	switch i {
		{{range $key, $val := $.Map.Data }}
			case {{$key}}:
				if count < {{$val.Length}} && count >= 0  {
					return {{"[]"}}{{$.Map.ValType}}{{"{"}}{{$val.Join}}{{"}"}}[count]
				}
		{{end}}
	}
	return 0
}
`
