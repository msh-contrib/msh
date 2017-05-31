package scope

type any interface{}

// Scope defines isolated environment for module
type Scope struct {
	parent *Scope
	store  map[string]any
}

// New creates new scope struture
func New(parent *Scope) *Scope {
	return &Scope{
		parent: parent,
		store:  make(map[string]any),
	}
}

// Get allows to access certain property in evn store
func (s *Scope) Get(name string) interface{} {
	if _, ok := s.store[name]; ok {
		return s.store[name]
	}

	if s.parent != nil {
		return s.parent.Get(name)
	}

	return nil
}

// Set allows to add properties to env store
func (s *Scope) Set(name string, value any) interface{} {
	s.store[name] = value
	return value
}
