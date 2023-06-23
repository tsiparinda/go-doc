package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
)

func main() {
	fset := token.NewFileSet()
	projectRoot := "../../vedicsociety/brucheion-pro"                               // replace with your project's root directory
	output, _ := os.Create("../../vedicsociety/brucheion-pro/docs/go-comments.txt") // output file for comments
	defer output.Close()

	err := filepath.Walk(projectRoot, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() || filepath.Ext(path) != ".go" {
			return nil
		}

		file, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
		if err != nil {
			return err
		}

		ast.Inspect(file, func(n ast.Node) bool {
			switch t := n.(type) {
			case *ast.FuncDecl:
				if t.Doc != nil {
					output.WriteString(fmt.Sprintf("Folder: %s - File: %s - Function: %s - Comment: %s\n", 
						filepath.Dir(path), filepath.Base(path), t.Name.Name, t.Doc.Text()))
				}
			case *ast.TypeSpec:
				if t.Doc != nil {
					output.WriteString(fmt.Sprintf("Folder: %s - File: %s - Type: %s - Comment: %s\n", 
						filepath.Dir(path), filepath.Base(path), t.Name.Name, t.Doc.Text()))
				}
			}
			return true
		})

		return nil
	})

	if err != nil {
		fmt.Printf("error walking the path %v: %v\n", projectRoot, err)
		return
	}
}