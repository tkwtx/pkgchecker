package main

import (
	"github.com/tkwtx/fmtchecker"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(fmtchecker.Analyzer) }
