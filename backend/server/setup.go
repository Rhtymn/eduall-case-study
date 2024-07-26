package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ServerOpts struct {
	ErrorHandler gin.HandlerFunc
}

func Setup(opts ServerOpts) *gin.Engine {
	router := gin.New()
	router.ContextWithFallback = true

	router.Use(
		gin.Recovery(),
		opts.ErrorHandler,
	)

	apiV1Group := router.Group("/api/v1")
	apiV1Group.GET("/ping", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "pong")
	})

	return router
}
