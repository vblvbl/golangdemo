package routers
import (
	"net/http"

	"golangdemo/controllers"

)


type MyMux struct {
}

func (p *MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.RequestURI
	var resp string
	switch path {
	case "/hello":
		resp = "hello"
	case "/haha":
		resp = "===="
	default:
		resp = "/"
	}
	if resp == "/hello" {
		contollers.SayhelloName(w, r)
	}


	http.NotFound(w, r)
	return
}


func Init() {
	mux := &MyMux{}
	http.ListenAndServe("9090", mux)

}