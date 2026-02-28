package main

import (
	"golang.org/x/tools/go/analysis/unitchecker"

	"github.com/FYNJYserty/log-linter-go/loglinter"
)

func main() {
	unitchecker.Main(loglinter.Analyzer)
}
