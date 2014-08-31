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
	f.Println("r.URL.User.Username:", r.URL.User.Username)
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

type mymux struct {
}

func (p *mymux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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
}

func SetupCustomedServ() {
	mux := &mymux{}
	
	f.Println("customed server is running ...")
	err := http.ListenAndServe(":9090", mux)
	if err != nil {
		log.Fatal("ListenAndServe failed ", err)
	}
}
