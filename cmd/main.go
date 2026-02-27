package main

import (
	"github.com/FYNJYserty/log-linter-go/loglinter"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(loglinter.Analyzer)
}
