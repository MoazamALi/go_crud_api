package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Movie struct {
	Id       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}
type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movie []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movie)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range movie {
		if item.Id == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var mov Movie
	_ = json.NewDecoder(r.Body).Decode(&mov)
	movie = append(movie, mov)
	json.NewEncoder(w).Encode(mov)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movie {
		if item.Id == params["id"] {
			movie = append(movie[:index], movie[index+1:]...)
			var mov Movie
			_ = json.NewDecoder(r.Body).Decode(&mov)
			mov.Id = params["id"]
			movie = append(movie, mov)
			json.NewEncoder(w).Encode(mov)
			return
		}
	}
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movie {
		if item.Id == params["id"] {
			movie = append(movie[:index], movie[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movie)
}
func main() {

	movie = append(movie, Movie{Id: "1", Isbn: "438227", Title: "Movie One", Director: &Director{Firstname: "John", Lastname: "Doe"}})
	movie = append(movie, Movie{Id: "2", Isbn: "438228", Title: "Movie Two", Director: &Director{Firstname: "Strve", Lastname: "Doe"}})
	movie = append(movie, Movie{Id: "3", Isbn: "438229", Title: "Movie Three", Director: &Director{Firstname: "jacob", Lastname: "Doe"}})
	movie = append(movie, Movie{Id: "4", Isbn: "438230", Title: "Movie Four", Director: &Director{Firstname: "John", Lastname: "Doe"}})
	movie = append(movie, Movie{Id: "5", Isbn: "438232", Title: "Movie Five", Director: &Director{Firstname: "david", Lastname: "Doe"}})

	r := mux.NewRouter()
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")
	fmt.Println("server started on port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
