package server

import (
	"net/http"

	"github.com/Rhtymn/eduall-case-study/backend/handler"
	"github.com/gin-gonic/gin"
)

type ServerOpts struct {
	ProductHandler *handler.ProductHandler
	ErrorHandler   gin.HandlerFunc
	CorsHandler    gin.HandlerFunc
}

func Setup(opts ServerOpts) *gin.Engine {
	router := gin.New()
	router.ContextWithFallback = true

	router.Use(
		gin.Recovery(),
		opts.ErrorHandler,
		opts.CorsHandler,
	)

	apiV1Group := router.Group("/api/v1")
	apiV1Group.GET("/ping", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "pong")
	})

	productsGroup := apiV1Group.Group("/products")
	productsGroup.GET(".", opts.ProductHandler.GetAll)

	return router
}
