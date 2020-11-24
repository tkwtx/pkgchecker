package fmtchecker

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const doc = "fmtchecker is ..."

// Analyzer is ...
var Analyzer = &analysis.Analyzer{
	Name: "fmtchecker",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

func run(pass *analysis.Pass) (interface{}, error) {
	isp := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.GenDecl)(nil),
		(*ast.Ident)(nil),
		(*ast.FuncDecl)(nil),
		(*ast.FuncLit)(nil),
		(*ast.BasicLit)(nil),
		(*ast.FuncLit)(nil),
	}

	isp.Preorder(nodeFilter, func(n ast.Node) {
		//fmt.Println(n," type:", reflect.TypeOf(n))
		switch n := n.(type) {
		case *ast.Ident:
			if fmtCheck(n.Name) {
				pass.Report(analysis.Diagnostic{
					Pos:     n.Pos(),
					Message: "use!",
				})
			}
		}
	})

	return nil, nil
}

func fmtCheck(s string) bool {
	return s == "Println" || s == "Printf" || s == "Print"
}
