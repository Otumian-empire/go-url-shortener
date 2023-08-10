package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/otumian-empire/go-url-shortener/repository"
	"github.com/otumian-empire/go-url-shortener/util"
	"github.com/otumian-empire/go-url-shortener/web"
)

func main() {
	defer util.Recover()

	ENV_CONST, err := godotenv.Read()

	if util.IsNotNil(err) {
		util.FLog(util.SERVER_LOADING_CREDENTIALS_ERROR)
	}

	databaseCre := mysql.Config{
		User:                 ENV_CONST["DATABASE_USERNAME"],
		Passwd:               ENV_CONST["DATABASE_PASSWORD"],
		Net:                  "tcp", // is tcp by default
		Addr:                 fmt.Sprintf("%s:%s", ENV_CONST["DATABASE_HOST"], ENV_CONST["DATABASE_PORT"]),
		DBName:               ENV_CONST["DATABASE_NAME"],
		AllowNativePasswords: true,
	}

	store, err := repository.NewStore(ENV_CONST["DATABASE_DRIVER_NAME"], databaseCre.FormatDSN())

	if util.IsNotNil(err) {
		util.FLog(err)
	}

	// this handler here is a not a handler as defined in the NewHandler
	// it is the route on passed to the new handler that is returned
	handler := web.NewHandler(*store, gin.Default())

	http.ListenAndServe(fmt.Sprintf(":%v", ENV_CONST["SERVER_PORT"]), handler)
}
