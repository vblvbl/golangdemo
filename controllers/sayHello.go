package contollers
import (
	"net/http"


	"fmt"
)

func SayhelloName(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello myroute!")

}