package fmtchecker

import (
	"fmt"
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
		(*ast.GenDecl)(nil),
	}

	isp.Preorder(nodeFilter, func(n ast.Node) {
		switch n := n.(type) {
		// import取得
		case *ast.GenDecl:

			if result := getImport(n); result != nil {
				fmt.Println(result)
			}
			// func取得
		case *ast.ExprStmt:
			call, ok := n.X.(*ast.CallExpr)
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
		}

	})

	return nil, nil
}

func getImport(n *ast.GenDecl) (imports []string) {
	for _, spec := range n.Specs {
		importSpec, ok := spec.(*ast.ImportSpec)
		if !ok {
			return nil
		}
		if importSpec.Name != nil {
			// Case: use alias import
			alias := importSpec.Name.Name
			if alias == "_" {
				continue
			}
			imports = append(imports, alias)
		} else {
			// Case: don't use alias import
			imports = append(imports, importSpec.Path.Value)
		}
	}
	return
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
