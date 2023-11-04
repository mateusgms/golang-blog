package main

import (
	"net/http"
	
	http_handler "github.com/mateusgms/golang-blog/handlers/http"
)

func main() {
	r := http_handler.NewRouter()
	r.Configure()
	http.ListenAndServe(":8080", r)
}