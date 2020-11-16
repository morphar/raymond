package raymond

import "testing"

const (
	VERBOSE = false
)

//
// Helpers
//

func barHelper(options *Options) string { return "bar" }

func echoHelper(str string, nb int) string {
	result := ""
	for i := 0; i < nb; i++ {
		result += str
	}

	return result
}

func boolHelper(b bool) string {
	if b {
		return "yes it is"
	}

	return "absolutely not"
}

func gnakHelper(nb int) string {
	result := ""
	for i := 0; i < nb; i++ {
		result += "GnAK!"
	}

	return result
}

//
// Tests
//

var helperTests = []Test{
	{
		"simple helper",
		`{{foo}}`,
		nil, nil,
		map[string]interface{}{"foo": barHelper},
		nil,
		`bar`,
	},
	{
		"helper with literal string param",
		`{{echo "foo" 1}}`,
		nil, nil,
		map[string]interface{}{"echo": echoHelper},
		nil,
		`foo`,
	},
	{
		"helper with identifier param",
		`{{echo foo 1}}`,
		map[string]interface{}{"foo": "bar"},
		nil,
		map[string]interface{}{"echo": echoHelper},
		nil,
		`bar`,
	},
	{
		"helper with literal boolean param",
		`{{bool true}}`,
		nil, nil,
		map[string]interface{}{"bool": boolHelper},
		nil,
		`yes it is`,
	},
	{
		"helper with literal boolean param",
		`{{bool false}}`,
		nil, nil,
		map[string]interface{}{"bool": boolHelper},
		nil,
		`absolutely not`,
	},
	{
		"helper with literal boolean param",
		`{{gnak 5}}`,
		nil, nil,
		map[string]interface{}{"gnak": gnakHelper},
		nil,
		`GnAK!GnAK!GnAK!GnAK!GnAK!`,
	},
	{
		"helper with several parameters",
		`{{echo "GnAK!" 3}}`,
		nil, nil,
		map[string]interface{}{"echo": echoHelper},
		nil,
		`GnAK!GnAK!GnAK!`,
	},
	{
		"#if helper with true literal",
		`{{#if true}}YES MAN{{/if}}`,
		nil, nil, nil, nil,
		`YES MAN`,
	},
	{
		"#if helper with false literal",
		`{{#if false}}YES MAN{{/if}}`,
		nil, nil, nil, nil,
		``,
	},
	{
		"#if helper with truthy identifier",
		`{{#if ok}}YES MAN{{/if}}`,
		map[string]interface{}{"ok": true},
		nil, nil, nil,
		`YES MAN`,
	},
	{
		"#if helper with falsy identifier",
		`{{#if ok}}YES MAN{{/if}}`,
		map[string]interface{}{"ok": false},
		nil, nil, nil,
		``,
	},
	{
		"#unless helper with true literal",
		`{{#unless true}}YES MAN{{/unless}}`,
		nil, nil, nil, nil,
		``,
	},
	{
		"#unless helper with false literal",
		`{{#unless false}}YES MAN{{/unless}}`,
		nil, nil, nil, nil,
		`YES MAN`,
	},
	{
		"#unless helper with truthy identifier",
		`{{#unless ok}}YES MAN{{/unless}}`,
		map[string]interface{}{"ok": true},
		nil, nil, nil,
		``,
	},
	{
		"#unless helper with falsy identifier",
		`{{#unless ok}}YES MAN{{/unless}}`,
		map[string]interface{}{"ok": false},
		nil, nil, nil,
		`YES MAN`,
	},
	{
		"#equal helper with same string var",
		`{{#equal foo "bar"}}YES MAN{{/equal}}`,
		map[string]interface{}{"foo": "bar"},
		nil, nil, nil,
		`YES MAN`,
	},
	{
		"#equal helper with different string var",
		`{{#equal foo "baz"}}YES MAN{{/equal}}`,
		map[string]interface{}{"foo": "bar"},
		nil, nil, nil,
		``,
	},
	{
		"#equal helper with same string vars",
		`{{#equal foo bar}}YES MAN{{/equal}}`,
		map[string]interface{}{"foo": "baz", "bar": "baz"},
		nil, nil, nil,
		`YES MAN`,
	},
	{
		"#equal helper with different string vars",
		`{{#equal foo bar}}YES MAN{{/equal}}`,
		map[string]interface{}{"foo": "baz", "bar": "tag"},
		nil, nil, nil,
		``,
	},
	{
		"#equal helper with same integer var",
		`{{#equal foo 1}}YES MAN{{/equal}}`,
		map[string]interface{}{"foo": 1},
		nil, nil, nil,
		`YES MAN`,
	},
	{
		"#equal helper with different integer var",
		`{{#equal foo 0}}YES MAN{{/equal}}`,
		map[string]interface{}{"foo": 1},
		nil, nil, nil,
		``,
	},
	{
		"#equal helper inside HTML tag",
		`<option value="test" {{#equal value "test"}}selected{{/equal}}>Test</option>`,
		map[string]interface{}{"value": "test"},
		nil, nil, nil,
		`<option value="test" selected>Test</option>`,
	},
	{
		"#equal full example",
		`{{#equal foo "bar"}}foo is bar{{/equal}}
{{#equal foo baz}}foo is the same as baz{{/equal}}
{{#equal nb 0}}nothing{{/equal}}
{{#equal nb 1}}there is one{{/equal}}
{{#equal nb "1"}}everything is stringified before comparison{{/equal}}`,
		map[string]interface{}{
			"foo": "bar",
			"baz": "bar",
			"nb":  1,
		},
		nil, nil, nil,
		`foo is bar
foo is the same as baz

there is one
everything is stringified before comparison`,
	},
	{
		"#gt helper with string value",
		`{{#gt 5.3 4.9}}YES MAN{{/gt}}`,
		nil, nil, nil, nil,
		"YES MAN",
	},
	{
		"#gt helper with context values",
		`{{#gt foo bar}}YES MAN{{/gt}}`,
		map[string]interface{}{"foo": 1337, "bar": -520.123},
		nil, nil, nil,
		"YES MAN",
	},
	{
		"#lt helper with string value",
		`{{#lt 4.9 5.0}}YES MAN{{/lt}}`,
		nil, nil, nil, nil,
		"YES MAN",
	},
	{
		"#lt helper with context values",
		`{{#lt foo bar}}YES MAN{{/lt}}`,
		map[string]interface{}{"foo": -520.1133, "bar": 1337},
		nil, nil, nil,
		"YES MAN",
	},
	{
		"#gte helper with string value",
		`{{#gte 5.3 5.3}}YES MAN{{/gte}}`,
		nil, nil, nil, nil,
		"YES MAN",
	},
	{
		"#gte helper with context values",
		`{{#gte foo bar}}YES MAN{{/gte}}`,
		map[string]interface{}{"foo": 1337, "bar": -520.123},
		nil, nil, nil,
		"YES MAN",
	},
	{
		"#lte helper with string value",
		`{{#lte 5.0 5.0}}YES MAN{{/lte}}`,
		nil, nil, nil, nil,
		"YES MAN",
	},
	{
		"#lte helper with context values",
		`{{#lte foo bar}}YES MAN{{/lte}}`,
		map[string]interface{}{"foo": -520.1133, "bar": 1337},
		nil, nil, nil,
		"YES MAN",
	},
}

//
// Let's go
//

func TestHelper(t *testing.T) {
	t.Parallel()

	launchTests(t, helperTests)
}

func TestRemoveHelper(t *testing.T) {
	RegisterHelper("testremovehelper", func() string { return "" })
	if _, ok := helpers["testremovehelper"]; !ok {
		t.Error("Failed to register global helper")
	}

	RemoveHelper("testremovehelper")
	if _, ok := helpers["testremovehelper"]; ok {
		t.Error("Failed to remove global helper")
	}
}

//
// Fixes: https://github.com/aymerick/raymond/issues/2
//

type Author struct {
	FirstName string
	LastName  string
}

func TestHelperCtx(t *testing.T) {
	RegisterHelper("template", func(name string, options *Options) SafeString {
		context := options.Ctx()

		template := name + " - {{ firstName }} {{ lastName }}"
		result, _ := Render(template, context)

		return SafeString(result)
	})

	template := `By {{ template "namefile" }}`
	context := Author{"Alan", "Johnson"}

	result, _ := Render(template, context)
	if result != "By namefile - Alan Johnson" {
		t.Errorf("Failed to render template in helper: %q", result)
	}
}
