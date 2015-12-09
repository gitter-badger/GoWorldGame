package main

import (
	"github.com/codegangsta/negroni"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	RegisterHandlers(mux)

	n := negroni.Classic()
	n.UseHandler(mux)
	n.Run(":3500")
}
