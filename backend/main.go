package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Rhtymn/eduall-case-study/backend/config"
	"github.com/Rhtymn/eduall-case-study/backend/database"
	"github.com/Rhtymn/eduall-case-study/backend/handler"
	"github.com/Rhtymn/eduall-case-study/backend/middleware"
	"github.com/Rhtymn/eduall-case-study/backend/repository/postgre"
	"github.com/Rhtymn/eduall-case-study/backend/server"
	"github.com/Rhtymn/eduall-case-study/backend/service"
)

func main() {
	err := config.Init()
	if err != nil {
		fmt.Printf("Error env initialization : %s", err.Error())
		os.Exit(1)
	}

	conf, err := config.Load()
	if err != nil {
		fmt.Printf("Error load env : %s", err.Error())
		os.Exit(1)
	}

	db, err := database.ConnectPostgreDB(conf.PostgreURL)
	if err != nil {
		fmt.Printf("Error connecting to DB : %s", err.Error())
		os.Exit(1)
	}

	errorHandler := middleware.NewErrorHandler()

	productRepository := postgre.NewProductRepository(db)

	productSrv := service.NewProductService(service.ProductServiceOpts{
		Product: productRepository,
	})

	productHandler := handler.NewProductHandler(handler.ProductHandlerOpts{
		ProductSrv: productSrv,
		Domain:     "product",
	})

	router := server.Setup(server.ServerOpts{
		ProductHandler: productHandler,
		ErrorHandler:   errorHandler,
	})

	server := &http.Server{
		Addr:    conf.ServerAddr,
		Handler: router,
	}

	fmt.Printf("Starting Server...\n")
	fmt.Printf("Server running on port %s\n", conf.ServerAddr)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		fmt.Printf("Error running server: %v\n", err)
	}
}
