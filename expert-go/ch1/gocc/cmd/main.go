package main

import (
	"github.com/n-kurasawa/gocc"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(gocc.Analyzer)
}
