package main

import (
	"fmt"
	"net/http"

	"github.com/dev-ayaa/track-space/pkg/config"

)

const portNumber = ":8080"

var app config.AppConfig

func main() {

	fmt.Println("Starting localhost server")

	server := &http.Server{Addr: portNumber, Handler: Routes(&app)}
	server.ListenAndServe()
	// fmt.Println("Project architecture loading")
}
