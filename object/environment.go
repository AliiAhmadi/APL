package object

func NewEnvironment() *Environment {
	return &Environment{
		store: make(map[string]Object),
	}
}

type Environment struct {
	store map[string]Object
}

func (environment *Environment) Get(name string) (Object, bool) {
	obj, ok := environment.store[name]
	return obj, ok
}

func (environment *Environment) Set(name string, val Object) Object {
	environment.store[name] = val
	return val
}
