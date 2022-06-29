package main

import (
	"fmt"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/singlechecker"
)

var analyzer = &analysis.Analyzer{
	Name: "cohesion",
	Doc:  "checks cohesion",
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	fmt.Println("Hello, world!")
	return nil, nil
}

func main() {
	singlechecker.Main(analyzer)
}
