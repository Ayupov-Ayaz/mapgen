package analysis

import (
	"errors"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"strings"

	"github.com/ayupov-ayaz/mapgen/analysis/internal/helpers"

	"github.com/ayupov-ayaz/mapgen/analysis/internal"
)

const (
	commentExp = "map_gen:"
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

func parseComment(decl *ast.GenDecl) (string, bool) {
	if decl.Doc != nil {
		if len(decl.Doc.List) > 0 {
			for _, c := range decl.Doc.List {
				if strings.Contains(c.Text, commentExp) {
					str := strings.Replace(c.Text, commentExp, "", 1)
					return strings.Replace(str, "//", "", 1), true
				}
			}
		}
	}

	return "", false
}

func analysisFileByMap(mapData MapParams) ([]internal.Result, error) {
	result := internal.NewResult(mapData.CountType)

	results := make([]internal.Result, 0, 2)

	fSet := token.NewFileSet()
	f, err := parser.ParseFile(fSet, mapData.FilePath, nil, parser.ParseComments)
	if err != nil {
		return nil, fmt.Errorf("ParseFile failed: %w", err)
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
				comment, ok := parseComment(decl)
				if !ok {
					continue
				}

				result.StructName = comment

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

							md := internal.NewMapData(*mapKeyVal, v)
							result.SetMapData(md)

							imps := helpers.GetNeedImports(*mapKeyVal, imports)
							result.SetImports(imps)

							results = append(results, *result)
						}
					}
				}
			}
		}
	}

	return results, nil
}
