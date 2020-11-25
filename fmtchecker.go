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

type TargetFunc struct {
	pkgName string
	funcs   []string
}

var targetFunc = TargetFunc{
	pkgName: "fmt",
	funcs:   []string{"Println", "Printf", "Print"},
}

func run(pass *analysis.Pass) (interface{}, error) {
	isp := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.ExprStmt)(nil),
	}

	isp.Preorder(nodeFilter, func(n ast.Node) {
		expr, ok := n.(*ast.ExprStmt)
		if !ok {
			return
		}
		call, ok := expr.X.(*ast.CallExpr)
		if !ok {
			return
		}
		selector, ok := call.Fun.(*ast.SelectorExpr)
		if !ok {
			return
		}
		if targetFunc.checkFunc(selector) {
			pass.Report(analysis.Diagnostic{
				Pos:     n.Pos(),
				Message: "fmt package is used!",
			})
		}
	})

	return nil, nil
}

func (t *TargetFunc) checkFunc(expr ast.Expr) bool {
	n, ok := expr.(*ast.SelectorExpr)
	if !ok {
		return false
	}
	id, ok := n.X.(*ast.Ident)
	if !ok {
		return false
	}
	if id.Name == t.pkgName {
		for _, v := range t.funcs {
			if n.Sel.Name == v {
				return true
			}
		}
	}

	return false
}
