package astpain

import (
	"go/ast"
)

// GetDeferredFunctionName возвращает имя функции, вызов которой был отложен через defer,
// если входящий node является *ast.DeferStmt.
func GetDeferredFunctionName(node ast.Node) string {
	ds, ok := node.(*ast.DeferStmt)
	if !ok {
		return ""
	}
	switch ft := ds.Call.Fun.(type) {
	case *ast.SelectorExpr:
		name := "." + ft.Sel.Name
		for {
			switch x := ft.X.(type) {
			case *ast.SelectorExpr:
				name = "." + x.Sel.Name + name
				ft = x
			case *ast.Ident:
				return x.Name + name
			}
		}
	case *ast.FuncLit:
		return "anonymous func"
	case *ast.Ident:
		return ft.Name
	}
	return ""
}
