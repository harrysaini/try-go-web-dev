package main

import (
	"log"
	"net/http"

	"github.com/harrysaini/try-go-web-dev/012-posgres/controllers"
	"github.com/harrysaini/try-go-web-dev/012-posgres/lib"
	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	ds := lib.InitDataStores()

	booksController := controllers.NewBooksController(ds.BooksDataStore)

	router.GET("/books", booksController.GetAllBooks)
	router.POST("/books", booksController.InsertBook)
	router.GET("/books/:isbn", booksController.GetBook)
	router.PUT("/books/:isbn", booksController.UpdateBook)
	router.DELETE("/books/:isbn", booksController.DeleteBook)

	log.Fatalln(http.ListenAndServe(":8181", router))
}
