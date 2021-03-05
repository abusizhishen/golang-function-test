package main

import (
	"fmt"
	"go/parser"
	"go/token"
	"log"
)

var (
	id = 1
	name = "whw"
	age= 18
)
// main func
func main() {
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, "/Users/whw/go/src/github.com/abusizhishen/golang-function-b/ast/b.go", nil, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
	}


	fmt.Printf("%+v", node)
}
