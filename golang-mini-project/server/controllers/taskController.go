package controllers

import (
	"encoding/json"
	"golang-mini-project/server/database"
	"golang-mini-project/server/entity"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//Get all tasking data
func GetAllTasking(w http.ResponseWriter, r *http.Request) {
	var tasking []entity.Tasking
	database.Connector.Find(&tasking)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tasking)
}

//Get tasking by id
func GetTaskingById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	var tasking entity.Tasking
	database.Connector.First(&tasking, key)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasking)
}

//Create tasking
func CreateTask(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var tasking entity.Tasking
	json.Unmarshal(requestBody, &tasking)
	database.Connector.Save(&tasking)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tasking)
}

//Update tasking
func UpdateTaskingById(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var tasking entity.Tasking
	json.Unmarshal(requestBody, &tasking)
	database.Connector.Save(&tasking)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tasking)
}

//Delete tasking
func DeleteTaskingById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	var tasking entity.Tasking
	id, _ := strconv.ParseInt(key, 10, 64)
	database.Connector.Where("id = ?", id).Delete(&tasking)
	w.WriteHeader(http.StatusNoContent)
}
