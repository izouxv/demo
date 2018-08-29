package router

import (
	"github.com/julienschmidt/httprouter"
	"mynotes/blog/controller"
	"net/http"
	"time"
	"mynotes/blog/config"
)

var (
	HttpServer = &http.Server{
		ReadTimeout: 60 * time.Second,
		WriteTimeout: 60 * time.Second,
		IdleTimeout: 60 * time.Second,
		//TLSConfig		: tls.Config{},
	}
	blogRouter *httprouter.Router
)

func RegisterRouter() {
	blogRouter = httprouter.New()
	blogRouter.GET("/",controller.Index)
	HttpServer.Addr = config.Config.AppServer.Addr
	HttpServer.Handler = blogRouter
}
