package parser

import (
	"github.com/cmaster11/raymond/ast"
)

type OnContentFunction = func(text string) string

type onContentVisitor struct {
	onContent OnContentFunction
}

// newOnContentVisitor instantiates a new onContentVisitor
func newOnContentVisitor(
	onContent OnContentFunction,
) *onContentVisitor {
	return &onContentVisitor{
		onContent: onContent,
	}
}

// WARNING: It must be called only once on AST.
func processOnContents(node ast.Node, onContent OnContentFunction) {
	node.Accept(newOnContentVisitor(onContent))
}

//
// Visitor interface
//

func (v *onContentVisitor) VisitProgram(program *ast.Program) interface{} {
	body := program.Body
	for _, current := range body {
		current.Accept(v)
	}

	return nil
}

func (v *onContentVisitor) VisitBlock(block *ast.BlockStatement) interface{} {
	if block.Program != nil {
		block.Program.Accept(v)
	}

	if block.Inverse != nil {
		block.Inverse.Accept(v)
	}

	return nil
}

func (v *onContentVisitor) VisitContent(node *ast.ContentStatement) interface{} {
	node.Value = v.onContent(node.Value)
	return nil
}

// NOOP
func (v *onContentVisitor) VisitMustache(mustache *ast.MustacheStatement) interface{} { return nil }
func (v *onContentVisitor) VisitPartial(mustache *ast.PartialStatement) interface{}   { return nil }
func (v *onContentVisitor) VisitComment(mustache *ast.CommentStatement) interface{}   { return nil }
func (v *onContentVisitor) VisitExpression(node *ast.Expression) interface{}          { return nil }
func (v *onContentVisitor) VisitSubExpression(node *ast.SubExpression) interface{}    { return nil }
func (v *onContentVisitor) VisitPath(node *ast.PathExpression) interface{}            { return nil }
func (v *onContentVisitor) VisitString(node *ast.StringLiteral) interface{}           { return nil }
func (v *onContentVisitor) VisitBoolean(node *ast.BooleanLiteral) interface{}         { return nil }
func (v *onContentVisitor) VisitNumber(node *ast.NumberLiteral) interface{}           { return nil }
func (v *onContentVisitor) VisitHash(node *ast.Hash) interface{}                      { return nil }
func (v *onContentVisitor) VisitHashPair(node *ast.HashPair) interface{}              { return nil }
