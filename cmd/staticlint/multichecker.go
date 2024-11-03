package main

import (
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/multichecker"
	"honnef.co/go/tools/staticcheck"
)

func main() {
	var myChecks []*analysis.Analyzer // коллекция анализаторов

	// всеанализаторы SA из staticcheck. Ничего не проверяем, просто берем все
	for _, v := range staticcheck.Analyzers {
		myChecks = append(myChecks, v.Analyzer)
	}

	myChecks = append(myChecks, exitAnalyzerStruct)

	multichecker.Main(myChecks...)
}
