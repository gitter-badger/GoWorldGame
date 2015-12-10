package main

import (
	"github.com/codegangsta/negroni"
	"net/http"
)
var portnum = ":12500"

func main() {
	mux := http.NewServeMux()
	RegisterHandlers(mux)

	n := negroni.Classic()
	n.UseHandler(mux)
	n.Run(portnum)
}
