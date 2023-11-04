package http

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Router struct {
	router *gin.Engine
}

type Route struct {
	Path    string
	Handler gin.HandlerFunc
	Methods []string
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.router.ServeHTTP(w, req)
}

func NewRouter() *Router {
	return &Router{
		router: gin.Default(),
	}
}

func (r *Router) Configure() {
	r.configureHealth()
	r.configureUser()
	r.configureNotFound()
}

func (r *Router) configure(subrouter *gin.RouterGroup, routes []Route) {
	for _, route := range routes {
		subrouter.Handle(route.Methods[0], route.Path, logRequest(route.Handler, "INFO"))
	}
}

func (r *Router) configureHealth() {
	api := r.router.Group("/api/v1")
	r.configure(api, []Route{
		{
			Path:    "/health",
			Handler: HealthHandler,
			Methods: []string{"GET"},
		},
	})
}

func (r *Router) configureUser() {
	api := r.router.Group("/api/v1")
	users := api.Group("/users")
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
			Path:    "/:id",
			Handler: GetUserHandler,
			Methods: []string{"GET"},
		},
		{
			Path:    "/:id",
			Handler: UpdateUserHandler,
			Methods: []string{"PUT"},
		},
		{
			Path:    "/:id",
			Handler: DeleteUserHandler,
			Methods: []string{"DELETE"},
		},
	})
}

func logRequest(handler gin.HandlerFunc, level string) gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Printf("[%s] %s %s %s", level, c.Request.RemoteAddr, c.Request.Method, c.Request.URL)
		handler(c)
	}
}

func (r *Router) configureNotFound() {
	r.router.NoRoute(func(c *gin.Context) {
		log.Printf("[ERROR] %s %s %s - 404 Not Found", c.Request.RemoteAddr, c.Request.Method, c.Request.URL)
		c.JSON(http.StatusNotFound, gin.H{"message": "404 - Not Found"})
	})
}
