package backend

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/jangie/jerq/resources"
)

type Router struct {
	Port int
}

func setHeaders(w http.ResponseWriter, path string) {
	if path[len(path)-3:] == ".js" {
		w.Header().Set("Content-Type", "application/javascript")
	}
}

func serveStaticRoutes(w http.ResponseWriter, req *http.Request) {
	var path = req.URL.Path[1:]
	data, err := resources.Asset(path)
	if err != nil {
		http.Error(w, fmt.Sprintf("we were unable to locate %s", path), http.StatusBadGateway)
	} else {
		setHeaders(w, path)
		s := string(data)
		fmt.Fprintf(w, s)
	}
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if strings.HasPrefix(req.URL.Path, "/static") {
		serveStaticRoutes(w, req)
	} else {
		fmt.Fprintf(w, "This is the jerq backend running on :%d", r.Port)
	}
}
