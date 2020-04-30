package analysis

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"

	"github.com/ayupov-ayaz/mapgen/analysis/internal"
)

func analysisFileByMap(mapData MapParams) (*internal.Result, error) {
	result := internal.NewResult("")

	fSet := token.NewFileSet()
	f, err := parser.ParseFile(fSet, mapData.FilePath, nil, 0)
	if err != nil {
		return result, fmt.Errorf("ParseFile failed: %w", err)
	}

	for _, decl := range f.Decls {
		switch decl := decl.(type) {
		case *ast.GenDecl:
			switch decl.Tok {
			case token.VAR:
				for _, spec := range decl.Specs {
					vSpec, ok := spec.(*ast.ValueSpec)
					if !ok {
						continue
					}

					if len(vSpec.Values) > 0 {
						for _, sv := range vSpec.Values {
							lit, err := internal.CastCompositeLit(sv)
							if err != nil {
								return nil, err
							}

							astIdent, err := internal.CastAstIdent(lit.Type)
							if err != nil {
								return nil, err
							}

							if astIdent.Name == mapData.MapType {
								mapKeyType, mapValType, err := internal.GetKeyValueType(astIdent)
								if err != nil {
									return nil, err
								}

								v, err := internal.GetMapValues(lit)
								if err != nil {
									return nil, err
								}

								md := internal.NewMapData(mapData.MapType, mapKeyType, mapValType, v)
								result.SetMapData(md)
							}
						}
					}
				}
			}
		}
	}

	return result, nil
}
