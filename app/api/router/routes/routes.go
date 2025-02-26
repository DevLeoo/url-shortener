package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Routes struct {
	Uri                string
	Method             string
	Fn                 func(http.ResponseWriter, *http.Request)
	NeedAuthentication bool
}

func Configure(r *mux.Router) *mux.Router {
	route := shortnerRoutes

	r.HandleFunc(route.Uri, route.Fn).Methods(route.Method)

	return r
}
