package main

import (
	"net/http"

	"github.com/nugrohoarr/golang-mux-jwt-boilerplate/controllers/authcontroller"
	"github.com/nugrohoarr/golang-mux-jwt-boilerplate/controllers/productcontroller"
	"github.com/nugrohoarr/golang-mux-jwt-boilerplate/middlewares"
	"github.com/nugrohoarr/golang-mux-jwt-boilerplate/models"

	"github.com/gorilla/mux"
)

func main() {

	models.ConnectDatabase()

	router := mux.NewRouter()

	router.HandleFunc("/signin", authcontroller.SignIn).Methods("POST")
	router.HandleFunc("/signup", authcontroller.SignUp).Methods("POST")
	router.HandleFunc("/logout", authcontroller.LogOut).Methods("GET")

	api := router.PathPrefix("/api").Subrouter()
	api.HandleFunc("/product", productcontroller.Index).Methods("GET")
	api.Use(middlewares.AuthMiddleware)

	http.ListenAndServe(":8080", router)

}
