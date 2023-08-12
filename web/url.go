package web

import (
	"fmt"
	"net/http"

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
	return func(context *gin.Context) {
		// get the data from the request body
		var url entity.Url

		if err := context.BindJSON(&url); util.IsNotNil(err) {
			util.Log(err)
			context.JSON(response.FailureMessageResponse(err.Error()))
			return
		}

		OriginalUrlHash := util.CreateHash(url.Original)

		shortUrl := fmt.Sprintf("%s/%s", context.Request.Host, OriginalUrlHash)

		id, err := handler.store.UrlStore.CreateUrl(shortUrl, url.Original)

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
			util.Log(err)
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
			util.Log(err)
			context.JSON(response.FailureMessageResponse(err.Error()))
			return
		}

		url, err := handler.store.UrlStore.Url(id)

		if util.IsNotNil(err) {
			util.Log(err)
			context.JSON(response.FailureMessageResponse(err.Error()))
			return
		}

		context.JSON(response.SuccessResponse(util.URLS_READ_SUCCESSFULLY, url))
	}
}

func (handler *UrlHandler) DeleteById() gin.HandlerFunc {
	return func(context *gin.Context) {
		id, err := util.ToInt(context.Param("id"))

		if util.IsNotNil(err) {
			util.Log(err)
			context.JSON(response.FailureMessageResponse(err.Error()))
			return
		}

		if err := handler.store.UrlStore.DeleteUrl(id); util.IsNotNil(err) {
			util.Log(err)
			context.JSON(response.FailureMessageResponse(err.Error()))
			return
		}

		context.JSON(response.SuccessMessageResponse(util.URL_DELETED_SUCCESSFULLY))
	}
}

func (handler *UrlHandler) GerOriginalUrl() gin.HandlerFunc {
	return func(context *gin.Context) {
		OriginalUrlHash := context.Param("hash")

		if util.IsEmpty(OriginalUrlHash) {
			context.JSON(response.FailureMessageResponse(util.VALUE_INVALID))
			return
		}

		shortUrl := fmt.Sprintf("%s/%s", context.Request.Host, OriginalUrlHash)

		url, err := handler.store.UrlStore.OriginalUrl(shortUrl)

		if util.IsNotNil(err) {
			util.Log(err)
			context.JSON(response.FailureMessageResponse(err.Error()))
			return
		}
		util.Log("Redirecting to:", url.Original)
		context.Redirect(http.StatusMovedPermanently, url.Original)
	}
}
