package web

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/otumian-empire/go-url-shortener/entity"
	"github.com/otumian-empire/go-url-shortener/repository"
	"github.com/otumian-empire/go-url-shortener/response"
	"github.com/otumian-empire/go-url-shortener/util"
)

type UrlHandler struct {
	store repository.Store
}

func (handler *UrlHandler) Create() gin.HandlerFunc {
	util.Log("Create a short a url")
	return func(context *gin.Context) {
		// get the data from the request body
		var url entity.Url

		util.Log(url.Original)

		if err := context.BindJSON(&url); util.IsNotNil(err) {
			util.Log(err)
			context.JSON(response.FailureMessageResponse(err.Error()))
			return
		}

		shortUrl := util.CreateHash(url.Original)

		serverURL := fmt.Sprintf("%s/%s", context.Request.Host, shortUrl)
		util.Log(serverURL)

		id, err := handler.store.UrlStore.CreateUrl(serverURL, url.Original)

		if util.IsNotNil(err) {
			util.Log(err)
			context.JSON(response.FailureMessageResponse(err.Error()))
			return
		}

		context.JSON(response.SuccessResponse(util.URL_CREATED_SUCCESSFULLY, id))
	}
}

func (handler *UrlHandler) Read() gin.HandlerFunc {
	return func(context *gin.Context) {
		urls, err := handler.store.UrlStore.Urls()

		if util.IsNotNil(err) {
			context.JSON(response.FailureMessageResponse(err.Error()))
			return
		}

		context.JSON(response.SuccessResponse(util.URLS_READ_SUCCESSFULLY, urls))
	}
}

func (handler *UrlHandler) ReadById() gin.HandlerFunc {
	return func(context *gin.Context) {
		id, err := util.ToInt(context.Param("id"))

		if util.IsNotNil(err) {
			context.JSON(response.FailureMessageResponse(err.Error()))
			return
		}

		url, err := handler.store.UrlStore.Url(id)

		if util.IsNotNil(err) {
			context.JSON(response.FailureMessageResponse(err.Error()))
			return
		}

		context.JSON(response.SuccessResponse(util.URLS_READ_SUCCESSFULLY, url))
	}
}

func (handler *UrlHandler) DeleteById() gin.HandlerFunc {
	panic("Not Implemented")
}

func (handler *UrlHandler) GerOriginalUrl() gin.HandlerFunc {
	panic("Not Implemented")
}
