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
		switch n := n.(type) {
		case *ast.ExprStmt:
			switch n2 := n.X.(type) {
			case *ast.CallExpr:
				switch n3 := n2.Fun.(type) {
				case *ast.SelectorExpr:
					if targetFunc.checkFunc(n3) {
						pass.Report(analysis.Diagnostic{
							Pos:     n.Pos(),
							Message: "use!",
						})
					}
				}
			}
		}
	})

	return nil, nil
}

func (t *TargetFunc) checkFunc(expr ast.Expr) bool {
	switch n := expr.(type) {
	case *ast.SelectorExpr:
		if n.X.(*ast.Ident).Name == t.pkgName {
			for _, v := range t.funcs {
				if n.Sel.Name == v {
					return true
				}
			}
		}
	}
	return false
}
