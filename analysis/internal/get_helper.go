package internal

import (
	"errors"
	"go/ast"
)

func GetMapValues(cl *ast.CompositeLit) (map[string]MapValData, error) {
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

func GetArrayType(expr ast.Expr) (string, error) {
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

func GetAstIdentName(expr ast.Expr) (string, error) {
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

	arr, err := GetArrayType(expr)
	if err != nil {
		return "", err
	}

	return arr, nil
}

func GetKeyValueTypeFromIdent(ident *ast.Ident) (string, string, error) {
	typeSpec, err := CastTypeSpec(ident.Obj.Decl)
	if err != nil {
		return "", "", err
	}

	mapType, err := CastMapType(typeSpec.Type)
	if err != nil {
		return "", "", err
	}
	mapKeyData, err := GetAstIdentName(mapType.Key)
	if err != nil {
		return "", "", err
	}
	mapValData, err := GetMapVal(mapType.Value)
	if err != nil {
		return "", "", err
	}

	return mapKeyData, mapValData, nil
}

func GetKeyValueTypeFromMapType(mapType *ast.MapType) (string, string, error) {
	selector, err := GetSelectorExpr(mapType.Key)
	if err != nil {
		return "", "", err
	}

	selectorX, err := CastAstIdent(selector.X)
	if err != nil {
		return "", "", err
	}

	key := selectorX.Name + "." + selector.Sel.Name

	arrType, err := CastArrayType(mapType.Value)
	if err != nil {
		return "", "", err
	}

	valData, err := CastAstIdent(arrType.Elt)
	if err != nil {
		return "", "", err
	}

	val := valData.Name

	return key, val, nil
}
