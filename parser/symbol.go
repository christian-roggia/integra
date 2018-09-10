package parser

import (
	"fmt"
)

type SymbolType string

const (
	SymbolInvalid  SymbolType = "INVALID"
	SymbolEquation            = "EQUATION"
	SymbolVariable            = "VARIABLE"
)

type SymbolTable struct {
	Name     string                `json:"scope"`
	Symbols  map[string]SymbolType `json:"symbols"`
	Children []*SymbolTable        `json:"children"`
	Parent   *SymbolTable          `json:"-"`
	Depth    int32                 `json:"depth"`
}

func newSymbolTable(name string, depth int32) *SymbolTable {
	return &SymbolTable{
		Name:    name,
		Symbols: make(map[string]SymbolType),
		Depth:   depth,
	}
}

func (tbl *SymbolTable) newChild(name string) *SymbolTable {
	child := newSymbolTable(name, tbl.Depth+1)
	child.Parent = tbl
	tbl.Children = append(tbl.Children, child)

	return child
}

func (tbl *SymbolTable) initialize(name string, t SymbolType) error {
	if tbl.has(name) {
		return fmt.Errorf("variable redeclaration in the same scope")
	}
	if tbl.parentHas(name) {
		fmt.Printf("Warning: variable %s redeclared in inner scope, change name to avoid variable shadowing.\n", name)
	}

	tbl.Symbols[name] = t
	return nil
}

func (tbl *SymbolTable) assign(name string, t SymbolType) error {
	if !tbl.has(name) && !tbl.parentHas(name) {
		return fmt.Errorf("invalid variable assignment to uninitialized variable")
	}

	if tbl.findType(name) != t {
		return fmt.Errorf("invalid variable assignment to variable with different type")
	}
	return nil
}

func (tbl *SymbolTable) findType(name string) SymbolType {
	if tbl.has(name) {
		return tbl.Symbols[name]
	}

	if tbl.parentHas(name) {
		return tbl.Parent.findType(name)
	}

	return SymbolInvalid
}

func (tbl *SymbolTable) has(name string) bool {
	_, ok := tbl.Symbols[name]
	return ok
}

func (tbl *SymbolTable) parentHas(name string) bool {
	if tbl.Parent != nil {
		if tbl.Parent.has(name) {
			return true
		}

		return tbl.Parent.parentHas(name)
	}

	return false
}
