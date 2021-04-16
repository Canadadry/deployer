package object

func NewEnclosedEnvironment(outer *Environment) *Environment {
	env := NewEnvironment(outer.builtins)
	env.outer = outer
	return env
}

func NewEnvironment(builtins map[string]*Builtin) *Environment {
	s := make(map[string]Object)
	return &Environment{store: s, outer: nil, builtins: builtins}
}

type Environment struct {
	builtins map[string]*Builtin
	store    map[string]Object
	outer    *Environment
}

func (e *Environment) GetBuiltIn(name string) (*Builtin, bool) {
	b, ok := e.builtins[name]
	return b, ok
}

func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	if !ok && e.outer != nil {
		obj, ok = e.outer.Get(name)
	}
	return obj, ok
}

func (e *Environment) Set(name string, val Object) Object {
	e.store[name] = val
	return val
}
