package main

import (
	"app/pkg/monkey"
	"app/pkg/monkey/object"
	"fmt"
	"os"
)

var store = map[string]string{}
var task = map[string]*object.Function{}

func Task(args ...object.Object) object.Object {
	if len(args) != 2 {
		return &object.Error{Message: fmt.Sprintf("task should have only two parameters,got %d", len(args))}
	}
	if args[1].Type() != object.FUNCTION_OBJ {
		return &object.Error{Message: fmt.Sprintf("task should take a string and a function ,got %s", args[1].Type())}
	}
	fn := args[1].(*object.Function)
	if len(fn.Parameters) != 0 {
		return &object.Error{Message: fmt.Sprintf("task function argument should have no parameters ,got %d", len(fn.Parameters))}
	}
	task[args[0].Inspect()] = fn
	return object.NULL
}

func Set(args ...object.Object) object.Object {
	if len(args) != 2 {
		return &object.Error{Message: fmt.Sprintf("set should have only two parameters,got %d", len(args))}
	}
	store[args[0].Inspect()] = args[1].Inspect()
	return object.NULL
}

func Get(args ...object.Object) object.Object {
	if len(args) != 1 {
		return &object.Error{Message: fmt.Sprintf("get should have only one parameter ,got %d", len(args))}
	}
	str, ok := store[args[0].Inspect()]
	if !ok {
		return object.NULL
	}
	return &object.String{Value: str}
}

func main() {
	if err := run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(args []string) error {
	if len(args) <= 1 {
		return fmt.Errorf("Expect a file to run")
	}
	f, err := os.Open(args[1])
	if err != nil {
		return err
	}
	eval := monkey.New(map[string]object.BuiltinFunction{
		"set":  Set,
		"get":  Get,
		"task": Task,
	})
	result := eval.Eval(f)
	if result.Type() == object.ERROR_OBJ {
		return fmt.Errorf(result.Inspect())
	}
	return nil
}
