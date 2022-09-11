package routes

import (
	"github.com.gorilla/mux"
	"github.com/chibuoyimm/go-bookstore/pkg/controllers"
)

var RegisterBookStoreRoutes = func(router *mux.Router){
	router.HandleFunc("/book/", cont)
}