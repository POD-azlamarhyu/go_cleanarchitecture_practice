package middleware

import (
	"net/http"
)

func ComposeMiddlewares(middlewares ...func(http.Handler) http.Handler) func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		for i:= range middlewares {
			h = middlewares[len(middlewares)-1-i](h)
		}
		return h
	}
}