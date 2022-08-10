package main

import (
	"golang-mini-project/server/controllers"
	"golang-mini-project/server/database"
	"golang-mini-project/server/entity"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	initDB()
	log.Println("Starting the HTTP server on port 8090")

	router := mux.NewRouter().StrictSlash(true)
	initaliseHandlers(router)
	log.Fatal(http.ListenAndServe(":8090", router))
}

func initaliseHandlers(router *mux.Router) {
	router.HandleFunc("/get", controllers.GetAllTasking).Methods("GET")
	router.HandleFunc("/get/{id}", controllers.GetTaskingById).Methods("GET")
	router.HandleFunc("/create", controllers.CreateTask).Methods("POST")
	router.HandleFunc("/update/{id}", controllers.UpdateTaskingById).Methods("PUT")
	router.HandleFunc("/delete/{id}", controllers.DeleteTaskingById).Methods("DELETE")
}

func initDB() {
	config :=
		database.Config{
			ServerName: "localhost:3306",
			User:       "root",
			Password:   "",
			DB:         "golang_mini_project",
		}

	connectionString := database.GetConnectionString(config)
	err := database.Connect(connectionString)
	if err != nil {
		panic(err.Error())
	}
	database.Migrate(&entity.Tasking{})
}
