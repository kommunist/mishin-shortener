package main

import (
	"context"
	"fmt"
	"os/exec"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/multichecker"
	"honnef.co/go/tools/staticcheck"

	"golang.org/x/tools/go/analysis/passes/nilness"
	"golang.org/x/tools/go/analysis/passes/shadow"
	"golang.org/x/tools/go/analysis/passes/slog"
)

// Создайте свой multichecker, состоящий из:
//
// + стандартных статических анализаторов пакета golang.org/x/tools/go/analysis/passes;
// + всех анализаторов класса SA пакета staticcheck.io;
// не менее одного анализатора остальных классов пакета staticcheck.io;
// двух или более любых публичных анализаторов на ваш выбор.

func main() {
	ciLintStart()
	cleanArch()

	var myChecks []*analysis.Analyzer // коллекция анализаторов

	// стандартных статических анализаторов пакета golang.org/x/tools/go/analysis/passes
	myChecks = append(myChecks, shadow.Analyzer)
	myChecks = append(myChecks, slog.Analyzer)
	myChecks = append(myChecks, nilness.Analyzer)

	// всех анализаторов класса SA пакета staticcheck.io
	for _, v := range staticcheck.Analyzers {
		myChecks = append(myChecks, v.Analyzer)
	}
	myChecks = append(myChecks, exitAnalyzerStruct)

	multichecker.Main(myChecks...)
}

func ciLintStart() {
	cmd := exec.CommandContext(
		context.Background(),
		"golangci-lint", "run", "./...",
	) // пока просто пусть будет Background, но надо сделать с Deadline, наверное
	out, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Printf("Ошибка выполнения: %s \n", err)
		fmt.Printf("Вывод линтера: %s \n", string(out))
		return
	}
}

func cleanArch() {
	cmd := exec.CommandContext(
		context.Background(),
		"go-cleanarch",
	) // пока просто пусть будет Background, но надо сделать с Deadline, наверное
	out, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Printf("Ошибка выполнения: %s \n", err)
		fmt.Printf("Вывод линтера: %s \n", string(out))
		return
	}
}
