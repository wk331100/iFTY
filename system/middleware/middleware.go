package middleware

type Middleware struct {
	Map []Map
}

type Map struct {
	Name string
	Function interface{}
}

func (m *Middleware) Register(name string, function interface{}){
	filed := Map{
		Name:     name,
		Function: function,
	}
	m.Map = append(m.Map, filed)
}
