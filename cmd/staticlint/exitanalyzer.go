package main

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
)

var exitAnalyzerStruct = &analysis.Analyzer{
	Name: "osExitAnalyzer",
	Doc:  "проверяет os.Exit в main функции main пакета",
	Run:  exitAnalyzer,
}

func exitAnalyzer(pass *analysis.Pass) (any, error) {
	for _, file := range pass.Files {
		if file.Name.Name != "main" { // если пакет не main, то бежим дальше
			continue
		}

		for _, anyDecl := range file.Decls {
			fD, ok := anyDecl.(*ast.FuncDecl)
			if !ok || fD.Name.Name != "main" { // если имя функции не main, то бежим дальше
				continue
			}

			// теперь мы находимся в main файле и в main функции
			ast.Inspect(fD, func(n ast.Node) bool {
				ce, ok := n.(*ast.CallExpr)
				if !ok { // если инспектируемое - это не вызов, то бежим дальше
					return true
				}

				se, ok := ce.Fun.(*ast.SelectorExpr) // проверяем, что это вызов метода на пакете
				if !ok {
					return true
				}

				ident, ok := se.X.(*ast.Ident) // Получаем идентификатор того, на чем вызвали
				if !ok {
					return true
				}

				if ident.Name == "os" && se.Sel.Name == "Exit" {
					pass.Reportf(ce.Pos(), "stop use os.Exit in main function main package")
				}

				return true
			})

		}

	}
	return nil, nil
}
