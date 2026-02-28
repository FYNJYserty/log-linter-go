package main

import (
	"golang.org/x/tools/go/analysis"

	"github.com/FYNJYserty/log-linter-go/loglinter"
)

var AnalyzerPlugin = map[string]*analysis.Analyzer{
	"loglinter": loglinter.Analyzer,
}
