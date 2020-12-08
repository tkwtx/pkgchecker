package main

import (
	"github.com/tkwtx/fmtchecker"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() { singlechecker.Main(fmtchecker.Analyzer) }
