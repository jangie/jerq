package backend

import (
	"fmt"
	"net/http"
)

type Router struct {
	Port int
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "This is the jerq backend running on :%d", r.Port)
}
