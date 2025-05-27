package routes

import (
	"github.com/gorilla/mux"
	"github.com/ujjavalparmar/go-bookstore/pkg/controllers"
)

var RegisterBookStoreRoutes = func(router *mux.Router){
	router.HandlerFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandlerFunc("/book/", controllers.GetBook).Methods("GET")
	router.HandlerFunc("/book/{bookid}", controllers.GetBookById).Methods("GET")
	router.HandlerFunc("/book/{bookid}", controllers.UpdateBook).Methods("PUT")
	router.HandlerFunc("/book/{bookid}", controllers.DeleteBook).Methods("DELETE")
}

