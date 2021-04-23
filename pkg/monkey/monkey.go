package monkey

import (
	"app/pkg/monkey/evaluator"
	"app/pkg/monkey/lexer"
	"app/pkg/monkey/object"
	"app/pkg/monkey/parser"
	"fmt"
	"io"
	"io/ioutil"
	"strings"
)

type Evalulator interface {
	Eval(io.Reader) object.Object
	SetEnv(name string, value object.Object)
}

type eval struct {
	env *object.Environment
}

func New(fn map[string]object.BuiltinFunction) *eval {
	env := object.NewEnvironment()
	for k, v := range evaluator.Builtins {
		env.Set(k, v)
	}
	for k, v := range fn {
		env.Set(k, &object.Builtin{Fn: v})
	}
	return &eval{env: env}
}

func (e *eval) Eval(cmd io.Reader) object.Object {
	content, err := ioutil.ReadAll(cmd)
	if err != nil {
		return &object.Error{Message: fmt.Sprintf("Can't read in put %v\n", err)}
	}

	l := lexer.New(string(content))
	p := parser.New(l)

	program := p.ParseProgram()
	if len(p.Errors()) != 0 {
		msg := strings.Join(p.Errors(), "\n")
		return &object.Error{Message: fmt.Sprintf("Parse errors : \n%s", msg)}
	}

	result := evaluator.Eval(program, e.env)
	if result != nil {
		if result.Type() == object.ERROR_OBJ {
			return &object.Error{Message: fmt.Sprintf("Runtime error : %s", result.Inspect())}
		}
	}
	return object.NULL
}

func (e *eval) SetEnv(name string, value object.Object) {
	e.env.Set(name, value)
}
