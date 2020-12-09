package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"github.com/gorilla/mux"
)

func handleRequests() {
	port := os.Getenv("PORT");
	if port == "" {
		port = "8080"
	}

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", hello)
	myRouter.HandleFunc("/status", status)
	myRouter.HandleFunc("/categories", categoryList)
	myRouter.HandleFunc("/category/{id}", categoryView)
	myRouter.HandleFunc("/category", categoryCreate).Methods("POST")

	fmt.Println("Starting up on " + port)
	log.Fatal(http.ListenAndServe(":" + port, myRouter))
}

func main() {
	Categories = []Category{
		{Id: "1", Title: "Food", Description: "First category"},
		{Id: "2", Title: "Gas", Description: "Fuel for our car"},
	}
	handleRequests()
}

type Category struct {
	Id string `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
}

var Categories []Category

func categoryList(w http.ResponseWriter, req *http.Request)  {
	json.NewEncoder(w).Encode(Categories)
}

func categoryView(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	key := vars["id"]
	for _, category := range Categories {
		if category.Id == key {
			json.NewEncoder(w).Encode(category)
		}
	}
}

func categoryCreate(w http.ResponseWriter, req *http.Request) {
	reqBody, _ := ioutil.ReadAll(req.Body)
	var category Category
	json.Unmarshal(reqBody, &category)
	// update our global Articles array to include
	// our new Article
	Categories = append(Categories, category)
	json.NewEncoder(w).Encode(Categories)
}

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Hello world!")
}

func status(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "STATUS: OK!")
}

