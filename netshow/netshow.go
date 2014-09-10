package netshow

import (
	f "fmt"
	"net/http"
	"strings"
	"log"
)

func sayHelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	f.Println("r.Form: ", r.Form)
	f.Println("path: ", r.URL.Path)
	
	f.Println("schema: ", r.URL.Scheme)
	f.Println("useragent: ", r.UserAgent())
	
	f.Println("r.Form[url_long]: ", r.Form["url_long"])
	
	for k, v := range r.Form {
		f.Println("key: ", k)
		f.Println("value: ", strings.Join(v, " "))
	}
	
	f.Fprintf(w, "Hello, liuyilin!")
}

func sayNext(w http.ResponseWriter, r *http.Request) {
	f.Fprintln(w, "welcome to the NEXT level!")
}

func showRequestInfo(w http.ResponseWriter, r *http.Request) {
	f.Println("r.cookie: ", r.Cookies())
	f.Println("r.ContentLength: ", r.ContentLength)
	f.Println("r.Host: ", r.Host)
	f.Println("r.method: ", r.Method)
	f.Println("r.close: ", r.Close)
	f.Println("r.proto: ", r.Proto)
	f.Println("r.ProtoMajor: ", r.ProtoMajor)
	f.Println("r.ProtoMinor: ", r.ProtoMinor)
	f.Println("r.RequestURI: ", r.RequestURI)
	f.Println("r.TLS", r.TLS)
	f.Println("r.RemoteAddr: ", r.RemoteAddr)
	f.Println("r.Trailer: ", r.Trailer)
	f.Println("r.TransferEncoding: ", r.TransferEncoding)
	f.Println("r.URL.schema: ", r.URL.Scheme)
	//username := r.URL.User.Username()
	//p, _ := r.URL.User.Password()
	//f.Println("r.URL.User.Username: ", username)
	//f.Println("r.URL.User.Username: ", p)
}

func SetupServ() {
	http.HandleFunc("/", sayHelloName)
	http.HandleFunc("/next", sayNext)
	
	f.Println("Serv is running...")
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListernAndServe: ", err)
	}
}

type State int

const (
	UNKNOWN State = iota
	GET_CATALOG
	PROVISION
	BINDING
	UNBINDING
	UNPROVISION
)

func(s State) String() string {
	switch s {
	case GET_CATALOG:
		return "GetCatalog"
	case PROVISION:
		return "Provision"
	case BINDING:
	 	return "Binding"
	case UNBINDING:
		return "Unbinding"
	case UNPROVISION:
		return "Unprovision"
	}
	
	return "Unknow Schema"
}

const (
	HTTP_METHOD_GET = "get"
	HTTP_METHOD_PUT = "put"
	HTTP_METHOD_POST = "post"
	HTTP_METHOD_DELETE = "delete"
)

const (
	PATTERN_SERVICE_INSTANCES = "service_instances"
	PATTERN_SERVICE_BINDINGS = "service_bindings"
)

//根据当前的path确定路由规则
//返回step表示目前采用什么方法
///v2/service_instances/:service-guid-id
///v2/service_instances/:instance_id/service_bindings/:id
//如果仅仅contain判断是否包含子串，这里都包含service_instance
//因此调整顺序，先判断是否包含service_bindings，如果有break，不判断service_instances
//这种方法有点tricky，不易维护，后面要修改掉，去掉对顺序的依赖

func route_dispatch(path string, method string) State{
	//如果是get方法，必然是获取catalog；如果是put/post/delete则再细分
	
	step := UNKNOWN
	lowerMethod := strings.ToLower(method)
	if lowerMethod == "get" {
		step = GET_CATALOG
	}
	
	patterns := []string {PATTERN_SERVICE_BINDINGS, PATTERN_SERVICE_INSTANCES}
	p := ""
	//必须用p将sp赋值出来，原因还未知！
	for _, sp := range(patterns) {
		b := strings.Contains(path, sp)
		//hit
		if b {
			p = sp
			break
		}
	}
	
	switch p {
	case PATTERN_SERVICE_BINDINGS:
		if lowerMethod == HTTP_METHOD_PUT || lowerMethod == HTTP_METHOD_POST {
			step = BINDING	
		} else if lowerMethod == HTTP_METHOD_DELETE {
			step = UNBINDING
		}
	//service-instance
	case PATTERN_SERVICE_INSTANCES:
		if lowerMethod == HTTP_METHOD_PUT || lowerMethod == HTTP_METHOD_POST {
			step = PROVISION
		} else if lowerMethod == "delete" {
			step = UNPROVISION
		}	
	}
	
	return step
}

type mymux struct {
}

func (p *mymux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	
	// get URL.PATH
	currPath := r.URL.Path
	method := r.Method
	f.Fprintf(w, "current path is: %s, method is %s\n", currPath, method)
	
	
	step := route_dispatch(currPath, method)
	/*
	if r.URL.Path == "/name" {
		sayHelloName(w, r)
	} else if r.URL.Path == "/next" {
		sayNext(w, r)
		f.Fprintf(w, "show Request Info")
		showRequestInfo(w, r)
	} else {
		f.Println("Not Found!")
		
		http.NotFound(w, r)
	}
	*/
	
	//根据step设置回调函数
	switch step {
		case GET_CATALOG:
		f.Println(step)
		case PROVISION:
		f.Println(step)
		case BINDING:
		f.Println(step)
		case UNBINDING:
		f.Println(step)
		case UNPROVISION:
		f.Println(step)
	}
}

func SetupCustomedServ() {
	mux := &mymux{}
	
	f.Println("customed server is running ...")
	err := http.ListenAndServe(":9090", mux)
	if err != nil {
		log.Fatal("ListenAndServe failed ", err)
	}
}
