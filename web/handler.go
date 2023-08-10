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

	// h.Get("/", h.ReadUrls())
	// h.Get("/:id", .Vote())
	// h.Get("/register", users.Register())
	// h.Post("/register", users.RegisterSubmit())
	// h.Get("/login", users.Login())
	// h.Post("/login", users.LoginSubmit())
	// h.Get("/logout", users.Logout())

	return h.router
}
