package main

import (
	"nestcheck"

	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(nestcheck.Analyzer) }
