package routers
import (
	"net/http"

	"golangdemo/controllers"
)


type MyMux struct {
}



func (p *MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		contollers.SayhelloName(w,r)
		return
	}
	http.NotFound(w, r)
	return
}




func Init() {
	mux := &MyMux{}
	http.ListenAndServe(":8080", mux)

}