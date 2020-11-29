package fmtchecker

import (
	"go/ast"
	"strings"

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
	packages []string
}

type pkg struct {
	packageName string
	funcName    string
	ok          bool
}

var targetFunc = new(TargetFunc)

func run(pass *analysis.Pass) (interface{}, error) {
	isp := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.ExprStmt)(nil),
		(*ast.GenDecl)(nil),
	}

	isp.Preorder(nodeFilter, func(n ast.Node) {
		switch n := n.(type) {
		case *ast.GenDecl:
			if result := getImport(n); result != nil {
				targetFunc.packages = result
			}
		case *ast.ExprStmt:
			call, ok := n.X.(*ast.CallExpr)
			if !ok {
				return
			}
			selector, ok := call.Fun.(*ast.SelectorExpr)
			if !ok {
				return
			}
			if result := targetFunc.checkFunc(selector); result.ok {
				pass.Report(analysis.Diagnostic{
					Pos:     n.Pos(),
					Message: "use " + result.packageName + "." + result.funcName,
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
			replaceStr := strings.Replace(importSpec.Path.Value, "\"", "", 2)
			imports = append(imports, replaceStr)
		}
	}
	return
}

func (t *TargetFunc) checkFunc(expr ast.Expr) pkg {
	n, ok := expr.(*ast.SelectorExpr)
	if !ok {
		return pkg{ok: false}
	}
	id, ok := n.X.(*ast.Ident)
	if !ok {
		return pkg{ok: false}
	}
	for _, v := range t.packages {
		if v == id.Name {
			return pkg{
				packageName: id.Name,
				funcName:    n.Sel.Name,
				ok:          true,
			}
		}
	}

	return pkg{ok: false}
}
