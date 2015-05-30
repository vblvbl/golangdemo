package routers
import (
	"net/http"


	"golangdemo/controllers"
	"log"
)


type MyMux struct {
}

func (p *MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.RequestURI
	switch path {
	case "/hello":
		controllers.SayhelloName(w, r)
	case "/handlecar":
		controllers.Hadlecar(w, r)
	default:
		http.NotFound(w, r)
	}

	return
}

func Init() {
	log.Println("服务器器启动成功:","8080端口")
	mux := &MyMux{}
	http.ListenAndServe(":8080", mux)

}