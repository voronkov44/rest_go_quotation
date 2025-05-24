package main

import (
	"fmt"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
	_ "rest_go_quotation_book/docs"
	"rest_go_quotation_book/internal/quotation"
)

// @title Quotation Book API
// @version 1.0
// @description API для работы с цитатами. Позволяет создавать, получать, удалять цитаты.

// @contact.name Andrew Voronkov
// @contact.email voronkovworkemail@gmail.com

// @host localhost:8081
// @BasePath /
func main() {
	router := http.NewServeMux()
	quotation.NewQuotationHandler(router)

	// Swagger docs endpoint
	router.Handle("/swagger/", httpSwagger.WrapHandler)

	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	fmt.Println("Server started at port 8081")
	server.ListenAndServe()
}
