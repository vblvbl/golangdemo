package main


import (

	"golangdemo/routers"

)

func main() {
// // Configuration
	PORT := config.Get("WEBSERVER_PORT")

	// Get the router from router.go
	router := routers.GetRouter()

	//Run HTTP Server
	fmt.Println("Running WebServer on Port " + PORT)

	log.Fatal(http.ListenAndServe(":"+PORT, router))
}