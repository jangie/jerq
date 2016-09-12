package main

import (
	"fmt"
	"net/http"

	"github.com/jangie/jerq/backend"
)

type server struct {
	router *backend.Router
}

func (s *server) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	s.router.ServeHTTP(w, req)
}

func main() {
	var port = 8080
	var s = &server{router: &backend.Router{Port: port}}
	go http.ListenAndServe(fmt.Sprintf(":%d", port), s)
	fmt.Printf("Listening on %d", port)
	select {}
}
