package env

type Environment struct {
	store map[string]interface{}
}

func NewEnvironment() *Environment {
	return &Environment{store: make(map[string]interface{})}
}

func (e *Environment) Set(name string, value interface{}) {
	e.store[name] = value
}

func (e *Environment) Get(name string) (interface{}, bool) {
	value, ok := e.store[name]
	return value, ok
}
