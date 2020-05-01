package analysis

import (
	"errors"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"

	"github.com/ayupov-ayaz/mapgen/analysis/internal/helpers"

	"github.com/ayupov-ayaz/mapgen/analysis/internal"
)

func parseImports(d *ast.GenDecl) ([]string, error) {
	imports := make([]string, len(d.Specs))
	for i, imp := range d.Specs {
		importStr, err := internal.ParseImport(imp)
		if err != nil {
			return nil, err
		}

		imports[i] = importStr
	}

	return imports, nil
}

func analysisFileByMap(mapData MapParams) (*internal.Result, error) {
	result := internal.NewResult("")

	fSet := token.NewFileSet()
	f, err := parser.ParseFile(fSet, mapData.FilePath, nil, 0)
	if err != nil {
		return result, fmt.Errorf("ParseFile failed: %w", err)
	}

	var (
		mapKeyVal *internal.MapKeyValue
		imports   []string
	)

	for _, decl := range f.Decls {
		switch decl := decl.(type) {
		case *ast.GenDecl:
			switch decl.Tok {
			case token.IMPORT:
				imports, err = parseImports(decl)
				if err != nil {
					return nil, err
				}

			case token.VAR:
				for _, spec := range decl.Specs {
					vSpec, ok := spec.(*ast.ValueSpec)
					if !ok {
						continue
					}

					if len(vSpec.Values) > 0 {
						for _, sv := range vSpec.Values {
							lit, err := internal.CastCompositeLit(sv)
							if err != nil && errors.Is(err, internal.ErrCastFailed) {
								continue
							}

							astIdent, ok := lit.Type.(*ast.Ident)
							if ok {
								if astIdent.Name != mapData.MapType {
									continue
								}

								mapKeyVal, err = internal.ParseKeyValueTypeFromIdent(astIdent)
								if err != nil {
									return nil, err
								}
							} else {
								mapType, err := internal.CastMapType(lit.Type)
								if err != nil {
									return nil, err
								}

								mapKeyVal, err = internal.ParseKeyValueTypeFromMapType(mapType)
								if err != nil {
									return nil, err
								}
							}

							v, err := internal.ParseMapValues(lit)
							if err != nil {
								return nil, err
							}

							md := internal.NewMapData(mapData.MapType, *mapKeyVal, v)
							result.SetMapData(md)

							imps := helpers.GetNeedImports(*mapKeyVal, imports)
							result.SetImports(imps)
						}
					}
				}
			}
		}
	}

	return result, nil
}
