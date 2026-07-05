package server

import (
	"net/http"
	"cleanarchitecture-practice/src/interface/middleware"
)



var (
	logger func(hh http.Handler) http.Handler
)

func initMiddlewares(){
	logger = middleware.Logger
}