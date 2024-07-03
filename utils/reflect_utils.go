package utils

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"os"
	"path/filepath"
	"reflect"
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

	// Walk through the directory and subdirectories
	err := filepath.Walk(packageDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip non-Go files and directories
		if !info.IsDir() && filepath.Ext(path) != ".go" {
			return nil
		}

		// Parse each Go file
		fset := token.NewFileSet()
		pkgs, parseErr := parser.ParseDir(fset, filepath.Dir(path), nil, parser.AllErrors)
		if parseErr != nil {
			return fmt.Errorf("failed to parse directory %s: %w", filepath.Dir(path), parseErr)
		}

		conf := types.Config{Importer: nil}
		info1 := &types.Info{
			Types: make(map[ast.Expr]types.TypeAndValue),
			Defs:  make(map[*ast.Ident]types.Object),
			Uses:  make(map[*ast.Ident]types.Object),
		}

		// Type check the package files
		for _, pkg := range pkgs {
			for _, file := range pkg.Files {
				_, typeErr := conf.Check(pkg.Name, fset, []*ast.File{file}, info1)
				if typeErr != nil {
					return fmt.Errorf("failed to type check file %s: %w", file.Name, typeErr)
				}

				// Find all struct types in the file
				for ident, obj := range info1.Defs {
					if obj, ok := obj.(*types.TypeName); ok {
						if _, ok := obj.Type().Underlying().(*types.Struct); ok {
							structTypes = append(structTypes, ident.Name)
						}
					}
				}
			}
		}

		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to walk directory: %w", err)
	}

	return structTypes, nil
}

//func ExtractStructTypes(packageDir string) ([]string, error) {
//	var structTypes []string
//
//	// Parse the package
//	fset := token.NewFileSet()
//	pkgs, err := parser.ParseDir(fset, packageDir, nil, parser.AllErrors)
//	if err != nil {
//		return nil, fmt.Errorf("failed to parse package: %w", err)
//	}
//
//	for _, pkg := range pkgs {
//		conf := types.Config{Importer: nil}
//		info := &types.Info{
//			Types: make(map[ast.Expr]types.TypeAndValue),
//			Defs:  make(map[*ast.Ident]types.Object),
//			Uses:  make(map[*ast.Ident]types.Object),
//		}
//
//		// Check each file in the package
//		for _, file := range pkg.Files {
//			_, err := conf.Check(pkg.Name, fset, []*ast.File{file}, info)
//			if err != nil {
//				log.Printf("failed to check package: %v", err)
//			}
//
//			// Find all struct types
//			for ident, obj := range info.Defs {
//				if obj, ok := obj.(*types.TypeName); ok {
//					if _, ok := obj.Type().Underlying().(*types.Struct); ok {
//						structTypes = append(structTypes, ident.Name)
//					}
//				}
//			}
//		}
//	}
//
//	return structTypes, nil
//}
