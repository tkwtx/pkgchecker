package main

import (
	"github.com/tkwtx/pkgchecker"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() { singlechecker.Main(pkgchecker.Analyzer) }
