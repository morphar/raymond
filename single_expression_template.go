package raymond

import (
	"errors"

	"github.com/aymerick/raymond/ast"
)

type SingleExpression struct {
	tpl  *Template
	expr *ast.Expression
}

// Parse instanciates a SingleExpression by parsing given source.
func ParseSingleExpression(source string) (*SingleExpression, error) {
	tpl := newTemplate(source)

	// parse template
	if err := tpl.parse(); err != nil {
		return nil, err
	}

	// validate single-expression-ness
	expr, err := validateSingleExpression(tpl.program.Body)
	if err != nil {
		return nil, err
	}

	return &SingleExpression{
		tpl:  tpl,
		expr: expr,
	}, nil
}

func validateSingleExpression(body []ast.Node) (*ast.Expression, error) {
	if len(body) == 0 {
		return nil, errors.New("the template is empty")
	}

	if len(body) != 1 {
		return nil, errors.New("you can only evaluate one expression at a time")
	}

	if body[0].Type() != ast.NodeMustache {
		return nil, errors.New("template does not contain a mustache statement")
	}

	stat, ok := body[0].(*ast.MustacheStatement)
	if !ok {
		return nil, errors.New("template does not contain a mustache statement")
	}

	return stat.Expression, nil
}

// EvalWith evaluates the single-expression template with given context and private data frame.
func (sexpr *SingleExpression) EvalWith(ctx interface{}, privData *DataFrame) (result interface{}) {
	// setup visitor
	v := newEvalVisitor(sexpr.tpl, ctx, privData)

	// visit AST
	return sexpr.expr.Accept(v)
}

// Eval evaluates template with given context.
func (sexpr *SingleExpression) Eval(ctx interface{}) (result interface{}) {
	return sexpr.EvalWith(ctx, nil)
}
