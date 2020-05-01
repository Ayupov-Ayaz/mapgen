package templates

const MapImpl = `//CODE GENERATED AUTOMATICALLY. DO NOT EDIT.
package {{.Package}}

type {{$.Type}} struct {}

func (p {{$.Type}}) Get(i {{$.Map.KV.Key.FullType}}, count int) {{$.Map.KV.Val.FullType}} {
	switch i {
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
