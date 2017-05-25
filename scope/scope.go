package scope

type any interface{}

type Scope struct {
	parent *Scope
	store  map[string]any
}

func New(parent *Scope) *Scope {
	return &Scope{
		parent: parent,
		store:  make(map[string]any),
	}
}

func (s *Scope) Get(name string) interface{} {
	if _, ok := s.store[name]; ok {
		return s.store[name]
	}

	if s.parent != nil {
		return s.parent.Get(name)
	}

	return nil
}

func (s *Scope) Set(name string, value any) any {
	s.store[name] = value
	return value
}
