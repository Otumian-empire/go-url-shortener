package web

import (
	"github.com/gin-gonic/gin"
	"github.com/otumian-empire/go-url-shortener/repository"
)

type Handler struct {
	store  repository.Store
	router *gin.Engine
}

func NewHandler(store repository.Store, router *gin.Engine) *gin.Engine {
	h := &Handler{
		store:  store,
		router: router,
	}

	// call the controllers and pass the repository
	urls := UrlHandler{store: store}
	// add other controllers here

	urlRoutes := h.router.Group("/url")
	{
		urlRoutes.GET("/", urls.Read())
		urlRoutes.GET("/:id", urls.ReadById())
		urlRoutes.POST("/", urls.Create())
	}

	return h.router
}
