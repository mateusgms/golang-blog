package http

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Router struct {
	router *mux.Router
}

type Route struct {
	Path    string
	Handler http.HandlerFunc
	Methods []string
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.router.ServeHTTP(w, req)
}

func NewRouter() *Router {
	return &Router{
		router: mux.NewRouter(),
	}
}

func (r *Router) Configure() {
	r.configureHealth()
	r.configureUser()
	r.configureNotFound()
}

func (r *Router) configure(subrouter *mux.Router, routes []Route) {
	for _, route := range routes {
		subrouter.HandleFunc(route.Path, logRequest(route.Handler, "INFO")).Methods(route.Methods...)
	}
}

func (r *Router) configureHealth() {
	api := r.router.PathPrefix("/api/v1").Subrouter()
	r.configure(api, []Route{
		{
			Path:    "/health",
			Handler: HealthHandler,
			Methods: []string{"GET"},
		},
	})
}

func (r *Router) configureUser() {
	api := r.router.PathPrefix("/api/v1").Subrouter()
	users := api.PathPrefix("/users").Subrouter()
	r.configure(users, []Route{
		{
			Path:    "",
			Handler: CreateUserHandler,
			Methods: []string{"POST"},
		},
		{
			Path:    "",
			Handler: ListUserHandler,
			Methods: []string{"GET"},
		},
		{
			Path:    "/{id}",
			Handler: GetUserHandler,
			Methods: []string{"GET"},
		},
		{
			Path:    "/{id}",
			Handler: UpdateUserHandler,
			Methods: []string{"PUT"},
		},
		{
			Path:    "/{id}",
			Handler: DeleteUserHandler,
			Methods: []string{"DELETE"},
		},
	})
}

func logRequest(handler http.HandlerFunc, level string) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		log.Printf("[%s] %s %s %s", level, req.RemoteAddr, req.Method, req.URL)
		handler(w, req)
	}
}
func (r *Router) configureNotFound() {
	r.router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		log.Printf("[ERROR] %s %s %s - 404 Not Found", req.RemoteAddr, req.Method, req.URL)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 - Not Found"))
	})
}
