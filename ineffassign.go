package main

import (
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
)

var (
	root            string
	dontRecurseFlag = flag.Bool("n", false, "don't recursively check paths")
)

func main() {
	flag.Parse()
	if len(flag.Args()) != 1 {
		fmt.Println("missing argument: filepath")
		return
	}

	var err error
	root, err = filepath.Abs(flag.Arg(0))
	if err != nil {
		fmt.Printf("Error finding absolute path :%s", err)
		return
	}

	filepath.Walk(root, checkPath)
}

func checkPath(path string, fi os.FileInfo, err error) error {
	if err != nil {
		fmt.Printf("Error during filesystem walk: %v\n", err)
		return nil
	}

	if fi.IsDir() {
		if *dontRecurseFlag && path != root {
			return filepath.SkipDir
		}
		return nil
	}
	if !strings.HasSuffix(path, ".go") {
		return nil
	}

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, path, nil, 0)
	if err != nil {
		return nil
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
	return nil
}

type checker struct {
	assignedNotUsed map[*ast.Object]*ast.Ident
	escapes         map[*ast.Object]bool
	loops, funcLits int
}

// TODO: Be less conservative: Take variable scope into consideration in loops and func literals.
func (chk *checker) Visit(n ast.Node) ast.Visitor {
	switch n := n.(type) {
	case *ast.AssignStmt:
		for _, x := range append(n.Rhs, n.Lhs...) {
			ast.Walk(chk, x)
		}
		if n.Tok == token.ASSIGN {
			for _, x := range n.Lhs {
				if i, ok := unparen(x).(*ast.Ident); ok {
					// Conservatively ignore assignments inside loops.
					if chk.loops == 0 && i.Obj != nil {
						chk.assignedNotUsed[i.Obj] = i
					}
				}
			}
		}
		return nil
	case *ast.BranchStmt:
		// A goto may act as a loop.  Conservatively ignore any assignments occurring before it.
		if n.Tok == token.GOTO {
			chk.assignedNotUsed = map[*ast.Object]*ast.Ident{}
		}
	case *ast.FuncType:
		if n.Results != nil {
			for _, f := range n.Results.List {
				for _, i := range f.Names {
					chk.escapes[i.Obj] = true
				}
			}
		}
	case *ast.Ident:
		delete(chk.assignedNotUsed, n.Obj)
		// Conservatively mark any variable mentioned in a func literal as escaping.
		if chk.funcLits > 0 {
			chk.escapes[n.Obj] = true
		}
	case *ast.UnaryExpr:
		if i, ok := unparen(n.X).(*ast.Ident); n.Op == token.AND && ok {
			chk.escapes[i.Obj] = true
		}
	case *ast.CallExpr:
		// A method call might implicitly take the address of its receiver, causing it to escape.
		// We can't do any better here without knowing the variable's type.
		if s, ok := unparen(n.Fun).(*ast.SelectorExpr); ok {
			if i, ok := unparen(s.X).(*ast.Ident); ok {
				chk.escapes[i.Obj] = true
			}
		}
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
