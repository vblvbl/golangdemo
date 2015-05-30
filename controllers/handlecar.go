package controllers
import (
	"net/http"
	"fmt"
)
func Hadlecar(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "this handlecar hello!")
}