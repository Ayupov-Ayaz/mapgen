package internal

import (
	"errors"
	"go/ast"
	"strings"
)

const (
	expComment = "//map_gen:"
)

func ParseMapValues(cl *ast.CompositeLit) (map[string]MapValData, error) {
	results := make(map[string]MapValData, len(cl.Elts))

	for _, v := range cl.Elts {
		kvExpr, err := CastKeyValueExpr(v)
		if err != nil {
			return nil, err
		}

		key, err := CastBasicLit(kvExpr.Key)
		if err != nil {
			return nil, err
		}

		clValues, err := CastCompositeLit(kvExpr.Value)
		if err != nil {
			return nil, err
		}

		values, err := CastSliceBasicList(clValues.Elts)
		if err != nil {
			return nil, err
		}

		vData := make([]string, len(values))
		for i := 0; i < len(values); i++ {
			vData[i] = values[i].Value
		}

		results[key.Value] = NewMapValData(vData)
	}

	return results, nil
}

func ParseArrayType(expr ast.Expr) (string, error) {
	arr, err := CastArrayType(expr)
	if err != nil {
		return "", err
	}

	ident, err := CastAstIdent(arr.Elt)
	if err != nil {
		return "", err
	}

	return ident.Name, nil
}

func ParseAstIdentName(expr ast.Expr) (string, error) {
	i, err := CastAstIdent(expr)
	if err != nil {
		return "", err
	}

	return i.Name, nil
}

func GetMapVal(expr ast.Expr) (string, error) {
	i, err := CastAstIdent(expr)
	if err != nil && !errors.Is(err, ErrCastFailed) {
		return "", err
	}

	if err == nil {
		return i.Name, nil
	}

	arr, err := ParseArrayType(expr)
	if err != nil {
		return "", err
	}

	return arr, nil
}

func ParseImport(spec ast.Spec) (string, error) {
	imp, err := CastImportSpec(spec)
	if err != nil {
		return "", err
	}

	return imp.Path.Value, nil
}

func ParseKeyValueTypeFromMapType(mapType *ast.MapType) (*MapKeyValue, error) {
	selector, err := GetSelectorExpr(mapType.Key)
	if err != nil {
		return nil, err
	}

	selectorX, err := CastAstIdent(selector.X)
	if err != nil {
		return nil, err
	}

	key := NewSpecType(selectorX.Name, selector.Sel.Name)

	arrType, err := CastArrayType(mapType.Value)
	if err != nil {
		return nil, err
	}

	valData, err := CastAstIdent(arrType.Elt)
	if err != nil {
		return nil, err
	}

	val := NewSpecType("", valData.Name)

	return NewMapKeyVal(key, val), nil
}

func ParseComment(decl *ast.GenDecl) (*Comment, error) {
	if decl.Doc != nil {
		if len(decl.Doc.List) > 0 {
			for _, c := range decl.Doc.List {
				if strings.Contains(c.Text, expComment) {
					c := NewComment(c.Text, expComment)
					if len(c.StructName) == 0 {
						return nil, errors.New("tag name not found")
					}

					return c, nil
				}
			}
		}
	}

	return nil, nil
}
