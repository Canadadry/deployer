package env

import (
	"app/internal/runner"
	"app/pkg/monkey/object"
	"fmt"
	"os"
)

type Task struct {
	Title       string
	Description string
	fn          *object.Function
}

type Environment struct {
	store map[string]string
	Tasks map[string]Task
}

func New() Environment {
	return Environment{
		store: map[string]string{},
		Tasks: map[string]Task{},
	}
}

func (e *Environment) AddTask(args ...object.Object) object.Object {
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
	e.Tasks[args[0].Inspect()] = Task{args[0].Inspect(), args[1].Inspect(), fn}
	return object.NULL
}

func (e *Environment) GetTask(args ...object.Object) object.Object {
	if len(args) != 1 {
		return &object.Error{Message: fmt.Sprintf("getTask should have only one parameter ,got %d", len(args))}
	}
	t, ok := e.Tasks[args[0].Inspect()]
	if !ok {
		return object.NULL
	}
	return t.fn
}

func (e *Environment) Set(args ...object.Object) object.Object {
	if len(args) != 2 {
		return &object.Error{Message: fmt.Sprintf("set should have only two parameters,got %d", len(args))}
	}
	e.store[args[0].Inspect()] = args[1].Inspect()
	return object.NULL
}

func (e *Environment) Get(args ...object.Object) object.Object {
	if len(args) != 1 {
		return &object.Error{Message: fmt.Sprintf("get should have only one parameter ,got %d", len(args))}
	}
	str, ok := e.store[args[0].Inspect()]
	if !ok {
		return object.NULL
	}
	return &object.String{Value: str}
}

func Env(args ...object.Object) object.Object {
	if len(args) != 1 {
		return &object.Error{Message: fmt.Sprintf("env should have only one parameter ,got %d", len(args))}
	}
	return &object.String{Value: os.Getenv(args[0].Inspect())}
}

func Run(r runner.Runner) func(args ...object.Object) object.Object {
	return func(args ...object.Object) object.Object {
		if len(args) != 1 {
			return &object.Error{Message: fmt.Sprintf("run should have only one parameters,got %d", len(args))}
		}
		return object.FromError(r.Run(args[0].Inspect()))
	}
}

func RunLocally(r runner.Runner) func(args ...object.Object) object.Object {
	return func(args ...object.Object) object.Object {
		if len(args) != 1 {
			return &object.Error{Message: fmt.Sprintf("runLocally should have only one parameters,got %d", len(args))}
		}
		return object.FromError(r.RunLocally(args[0].Inspect()))
	}
}

func Upload(r runner.Runner) func(args ...object.Object) object.Object {
	return func(args ...object.Object) object.Object {
		if len(args) != 2 {
			return &object.Error{Message: fmt.Sprintf("upload should have only two parameters,got %d", len(args))}
		}
		return object.FromError(r.Upload(args[0].Inspect(), args[1].Inspect()))
	}
}

func Download(r runner.Runner) func(args ...object.Object) object.Object {
	return func(args ...object.Object) object.Object {
		if len(args) != 2 {
			return &object.Error{Message: fmt.Sprintf("download should have only two parameters,got %d", len(args))}
		}
		return object.FromError(r.Download(args[0].Inspect(), args[1].Inspect()))
	}
}
