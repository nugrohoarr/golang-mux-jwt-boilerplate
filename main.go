package main

import (
	"net/http"

	"github.com/jeypc/go-jwt-mux/controllers/authcontroller"
	"github.com/jeypc/go-jwt-mux/models"

	"github.com/gorilla/mux"
)

func main() {

	models.ConnectDatabase()

	router := mux.NewRouter()

	router.HandleFunc("/signin", authcontroller.SignIn).Methods("POST")
	router.HandleFunc("/signup", authcontroller.SignUp).Methods("POST")
	router.HandleFunc("/logout", authcontroller.LogOut).Methods("GET")

	http.ListenAndServe(":8080", router)

}
