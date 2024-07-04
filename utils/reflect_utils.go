package utils

import (
	"fmt"
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"reflect"
	//"reflect"
)

func FillObjFieldValue(fieldString map[string]string, object interface{}) {
	// 获取接口的反射值
	val := reflect.ValueOf(object)

	for {
		if val.Kind() == reflect.Interface { // 是interface先取值
			val = val.Elem()
		} else if val.Kind() == reflect.Ptr { // 获取指针指向的值
			val = val.Elem()
		} else {
			break
		}
	}

	// 获取实际类型
	typ := val.Type()

	// 输出结构体的类型
	log.Info("obj type: ", typ.Name())

}

func CreateInstance(t reflect.Type) interface{} {
	return reflect.New(t).Elem().Interface()
}
func ExtractStructTypes(packageDir string) ([]string, error) {
	var structTypes []string

	// Parse the package
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, packageDir, nil, parser.AllErrors)
	if err != nil {
		return nil, fmt.Errorf("failed to parse package: %w", err)
	}

	for _, pkg := range pkgs {
		conf := types.Config{Importer: importer.Default()}
		info := &types.Info{
			Types: make(map[ast.Expr]types.TypeAndValue),
			Defs:  make(map[*ast.Ident]types.Object),
			Uses:  make(map[*ast.Ident]types.Object),
		}

		// Check each file in the package
		for _, file := range pkg.Files {
			_, err := conf.Check(pkg.Name, fset, []*ast.File{file}, info)
			if err != nil {
				log.Printf("failed to check package: %v", err)
			}

			// Find all struct types
			for ident, obj := range info.Defs {
				if obj, ok := obj.(*types.TypeName); ok {
					if _, ok := obj.Type().Underlying().(*types.Struct); ok {
						structTypes = append(structTypes, ident.Name)
					}
				}
			}
		}
	}

	return structTypes, nil
}
