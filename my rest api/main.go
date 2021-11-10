package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "3000"
)

type Recipe struct {
	ID   string `json:"id"`
	Isbn string `json:"isbn"`
	Food string `json:"food"`
	Chef *Chef  `json:"chef"`
}

//Chef struct
type Chef struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

var recipe []Recipe

func getRecipes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(recipe)
}

func getRecipe(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range recipe {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Recipe{})
}
func postRecipe(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var recipey Recipe
	_ = json.NewDecoder(r.Body).Decode(&recipey)
	recipey.ID = strconv.Itoa(rand.Intn(10000000))
	recipe = append(recipe, recipey)
	json.NewEncoder(w).Encode(recipey)
}
func putRecipe(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range recipe {
		if item.ID == params["id"] {
			recipe = append(recipe[:index], recipe[index+1:]...)

			var recipey Recipe
			_ = json.NewDecoder(r.Body).Decode(&recipey)
			recipey.ID = params["id"]
			recipe = append(recipe, recipey)
			json.NewEncoder(w).Encode(recipey)
			return
		}
	}
	json.NewEncoder(w).Encode(recipe)

}
func deleteRecipe(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range recipe {
		if item.ID == params["id"] {
			recipe = append(recipe[:index], recipe[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(recipe)
}

func main() {
	r := mux.NewRouter()

	recipe = append(recipe, Recipe{ID: "1", Isbn: "4566433", Food: "Recipe one", Chef: &Chef{FirstName: "John", LastName: "Doe"}})
	recipe = append(recipe, Recipe{ID: "2", Isbn: "4566456", Food: "Recipe two", Chef: &Chef{FirstName: "Jon", LastName: "Snow"}})

	r.HandleFunc("/recipe", getRecipes).Methods("GET")
	r.HandleFunc("/recipe/{id}", getRecipe).Methods("GET")
	r.HandleFunc("/recipe", postRecipe).Methods("POST")
	r.HandleFunc("/recipe/{id}", putRecipe).Methods("PUT")
	r.HandleFunc("/recipe/{id}", deleteRecipe).Methods("DELETE")
	err := http.ListenAndServe(CONN_HOST+":"+CONN_PORT, r)
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
