package server

import (
	"net/http"
	"cleanarchitecture-practice/src/interface/middleware"
)

func NewMux() *http.ServeMux {
	mux := http.NewServeMux()
	initControllers()
	initMiddlewares()
	registerRoutes(mux)
	return mux
}

func registerRoutes(mux *http.ServeMux){
	{
		mux.Handle("POST /signup", middleware.ComposeMiddlewares(logger)(signupController))
	}
}