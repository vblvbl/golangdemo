package routers
import (
	"net/http"

	"fmt"
	"log"
)

type MyMux struct {
}

func (p *MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/hello" {
		SayhelloName(w, r)
		return
	}
	http.NotFound(w, r)
	return
}
func SayhelloName(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello myroute!")
}
func Init() {
	log.Println("服务器启动：", "端口9090")
	mux := &MyMux{}
 http.ListenAndServe("9090", mux)
}