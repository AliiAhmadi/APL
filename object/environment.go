package object

func NewEnvironment() *Environment {
	return &Environment{
		store: make(map[string]Object),
		outer: nil,
	}
}

type Environment struct {
	store map[string]Object
	outer *Environment
}

func (environment *Environment) Get(name string) (Object, bool) {
	obj, ok := environment.store[name]
	if !ok && environment.outer != nil {
		obj, ok = environment.outer.Get(name)
	}
	return obj, ok
}

func (environment *Environment) Set(name string, val Object) Object {
	environment.store[name] = val
	return val
}

func NewEncloseEnvironment(outer *Environment) *Environment {
	env := NewEnvironment()
	env.outer = outer
	return env
}
