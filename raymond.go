// Package raymond provides handlebars evaluation
package raymond

// Render parses a template and renders it with given context
//
// Note that this function call is not optimal as your template is parsed everytime you call it. You should use Parse() function instead.
func Render(source string, ctx interface{}) (string, error) {
	// parse template
	tpl, err := Parse(source)
	if err != nil {
		return "", err
	}

	// renders template
	str, err := tpl.Exec(ctx)
	if err != nil {
		return "", err
	}

	return str, nil
}

// MustRender parses a template and evaluates it with given context. It panics on error.
//
// Note that this function call is not optimal as your template is parsed everytime you call it. You should use Parse() function instead.
func MustRender(source string, ctx interface{}) string {
	return MustParse(source).MustExec(ctx)
}

// EvalSingleExpression parses a template with one expression and evaluates it
// with given context.
//
// This is useful to evaluate helper calls etc.
func EvalSingleExpression(source string, ctx interface{}) (interface{}, error) {
	// parse template
	sexpr, err := ParseSingleExpression(source)
	if err != nil {
		return "", err
	}

	// evaluate expression
	return sexpr.Eval(ctx), nil
}
