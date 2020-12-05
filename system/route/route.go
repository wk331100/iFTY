package route

const METHOD_GET = "GET"
const METHOD_POST = "POST"
const METHOD_REQUEST = "REQUEST"

type Route struct {
	Prefix string
	Map []Map
}

type Map struct {
	Path string
	Function interface{}
	Method string
}

type FieldParam struct {
	Path string
	function interface{}
}


func (r *Route) Group(name string, f func()) {
	r.Prefix = name
	f()
	r.Prefix = ""
}


func (r *Route) Get(path string, function interface{}){
	r.setField(path, function, METHOD_GET)
}

func (r *Route) Post(path string, function interface{}){
	r.setField(path, function, METHOD_POST)
}

func (r *Route) Any(path string, function interface{}){
	r.setField(path, function, METHOD_REQUEST)
}


func (r *Route) setField(path string, function interface{}, method string){
	if r.Prefix != "" {
		path = r.Prefix + path
	}
	field := Map{
		Path:     path,
		Function: function,
		Method:   method,
	}
	r.Map = append(r.Map, field)
}