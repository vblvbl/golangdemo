package contollers
import (
	"net/http"
	"fmt"

	"godemo/routers"
)
type SayhelloName struct  {
	routers.MyMux
}

func (S *SayhelloName)SayhelloName(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello myroute!")

}