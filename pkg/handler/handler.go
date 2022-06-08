package handler

import (
	"fmt"
	"net/http"

)

func HomePage(wr http.ResponseWriter, rq *http.Request) {
	fmt.Println("Welcome to Track Space")
	fmt.Sprintln("Hello Everyone Welcome to Track Space")

}