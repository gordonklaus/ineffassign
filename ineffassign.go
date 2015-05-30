package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"strings"

	"github.com/kr/fs"
)

func main() {
	if len(os.Args) < 1 {
		fmt.Println("missing argument: filepath")
		return
	}

	walker := fs.Walk(os.Args[1])
	for walker.Step() {
		if err := walker.Err(); err != nil {
			fmt.Printf("Error during filesystem walk: %v\n", err)
			continue
		}

		if walker.Stat().IsDir() || !strings.HasSuffix(walker.Path(), ".go") {
			continue
		}

		fset := token.NewFileSet()
		f, err := parser.ParseFile(fset, walker.Path(), nil, 0)
		if err != nil {
			continue
		}

		chk := &checker{map[*ast.Object]*ast.Ident{}, map[*ast.Object]bool{}, 0, 0}
		for _, d := range f.Decls {
			if d, ok := d.(*ast.GenDecl); ok && d.Tok == token.VAR {
				for _, s := range d.Specs {
					for _, i := range s.(*ast.ValueSpec).Names {
						chk.escapes[i.Obj] = true
					}
				}
			}
		}
		ast.Walk(chk, f)
		for _, i := range chk.assignedNotUsed {
			if !chk.escapes[i.Obj] {
				fmt.Println(fset.Position(i.Pos()), i.Name)
			}
		}
	}
}

type checker struct {
	assignedNotUsed map[*ast.Object]*ast.Ident
	escapes         map[*ast.Object]bool
	loops, funcLits int
}

func (chk *checker) Visit(n ast.Node) ast.Visitor {
	switch n := n.(type) {
	case *ast.FuncType:
		if n.Results != nil {
			for _, f := range n.Results.List {
				for _, i := range f.Names {
					chk.escapes[i.Obj] = true
				}
			}
		}
	case *ast.AssignStmt:
		for _, x := range append(n.Rhs, n.Lhs...) {
			ast.Walk(chk, x)
		}
		if n.Tok == token.ASSIGN {
			for _, x := range n.Lhs {
				if i, ok := unparen(x).(*ast.Ident); ok {
					// TODO: ignore chk.loops if i.Obj was declared in the current loop
					if chk.loops == 0 && i.Obj != nil {
						chk.assignedNotUsed[i.Obj] = i
					}
				}
			}
		}
		return nil
	case *ast.ForStmt:
		walk(chk, n.Init)
		chk.loops++
		walk(chk, n.Cond)
		walk(chk, n.Post)
		walk(chk, n.Body)
		chk.loops--
		return nil
	case *ast.RangeStmt:
		walk(chk, n.X)
		chk.loops++
		walk(chk, n.Key)
		walk(chk, n.Value)
		walk(chk, n.Body)
		chk.loops--
		return nil
	case *ast.FuncLit:
		walk(chk, n.Type)
		chk.funcLits++
		walk(chk, n.Body)
		chk.funcLits--
		return nil
	case *ast.BranchStmt:
		if n.Tok == token.GOTO {
			// conservative
			chk.assignedNotUsed = map[*ast.Object]*ast.Ident{}
		}
	case *ast.Ident:
		delete(chk.assignedNotUsed, n.Obj)
		if chk.funcLits > 0 {
			chk.escapes[n.Obj] = true
		}
	case *ast.UnaryExpr:
		if i, ok := unparen(n.X).(*ast.Ident); n.Op == token.AND && ok {
			chk.escapes[i.Obj] = true
		}
	case *ast.CallExpr:
		if s, ok := unparen(n.Fun).(*ast.SelectorExpr); ok {
			if i, ok := unparen(s.X).(*ast.Ident); ok {
				chk.escapes[i.Obj] = true
			}
		}
	}
	return chk
}

func walk(v ast.Visitor, n ast.Node) {
	if n != nil {
		ast.Walk(v, n)
	}
}

func unparen(x ast.Expr) ast.Expr {
	if p, ok := x.(*ast.ParenExpr); ok {
		return unparen(p.X)
	}
	return x
}

type T int

func (T) f() {}

func f() {
	x := T(0)
	x = 1
	x.f()
	for x = 3; ; {
		for {
		}
		x := 0
		x = 1
		_ = x
		x = 2
	}
}
