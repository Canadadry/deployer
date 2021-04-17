package internal

import (
	"app/pkg/monkey/object"
	"fmt"
)

type Task struct {
	Title       string
	Description string
	fn          *object.Function
}

var Store = map[string]string{}
var Tasks = map[string]Task{}

func AddTask(args ...object.Object) object.Object {
	if len(args) != 3 {
		return &object.Error{Message: fmt.Sprintf("task should have three parameters,got %d\n usage : tast(title, description, function)", len(args))}
	}
	if args[2].Type() != object.FUNCTION_OBJ {
		return &object.Error{Message: fmt.Sprintf("task should take a string and a function ,got %s", args[1].Type())}
	}
	fn := args[2].(*object.Function)
	if len(fn.Parameters) != 0 {
		return &object.Error{Message: fmt.Sprintf("task function argument should have no parameters ,got %d", len(fn.Parameters))}
	}
	Tasks[args[0].Inspect()] = Task{args[0].Inspect(), args[1].Inspect(), fn}
	return object.NULL
}

func GetTask(args ...object.Object) object.Object {
	if len(args) != 1 {
		return &object.Error{Message: fmt.Sprintf("getTask should have only one parameter ,got %d", len(args))}
	}
	t, ok := Tasks[args[0].Inspect()]
	if !ok {
		return object.NULL
	}
	return t.fn
}

func Set(args ...object.Object) object.Object {
	if len(args) != 2 {
		return &object.Error{Message: fmt.Sprintf("set should have only two parameters,got %d", len(args))}
	}
	Store[args[0].Inspect()] = args[1].Inspect()
	return object.NULL
}

func Get(args ...object.Object) object.Object {
	if len(args) != 1 {
		return &object.Error{Message: fmt.Sprintf("get should have only one parameter ,got %d", len(args))}
	}
	str, ok := Store[args[0].Inspect()]
	if !ok {
		return object.NULL
	}
	return &object.String{Value: str}
}
