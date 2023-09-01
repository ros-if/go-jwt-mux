package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/roshif-study/task-5-vix-btpns--ROSHIF-/controllers/authcontroller"
	"github.com/roshif-study/task-5-vix-btpns--ROSHIF-/models"
)

func main() {

	models.ConnectDatabase()
	r := mux.NewRouter()

	r.HandleFunc("users/register", authcontroller.Register).Methods("POST")
	r.HandleFunc("users/login", authcontroller.Login).Methods("POST")
	r.HandleFunc("users/logout", authcontroller.Logout).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))

}