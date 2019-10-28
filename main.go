package main

import (
	"fmt"
	"encoding/json"
	"math/rand"
	"log"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

// Movie model
type Movie struct {
	ID string `json: "id"`
	Name string `json: "name"`
	Release string `json: "release"`
	Director *Director `json: "director"`
}

// Director model 
type Director struct {
	FirstName string `json: "firstname"`
	LastName string `json: "lastname"`
}

var movies = []Movie {
	Movie{
		ID: "1",
		Name: "The Dark Knight",
		Release: " 18 July 2008",
		Director: &Director{
			FirstName: "Christopher",
			LastName: "Nolan",
		},
	},
	Movie{
		ID: "2",
		Name: "Batman Begins",
		Release: "15 June 2005",
		Director: &Director{
			FirstName: "Christopher",
			LastName: "Nolan",
		},
	},
	Movie{
		ID: "3",
		Name: "The Dark Knight Rises",
		Release: "20 July 2012",
		Director: &Director{
			FirstName: "Christopher",
			LastName: "Nolan",
		},
	},
}

func addMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var movie Movie

	_ = json.NewDecoder(r.Body).Decode(&movie)

	movie.ID = strconv.Itoa(rand.Intn(1000000))
	movies = append(movies, movie)

	json.NewEncoder(w).Encode(movie)
}

func showMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// get params
	params := mux.Vars(r)

	for _, movie := range movies {
		if movie.ID == params["id"] {
			json.NewEncoder(w).Encode(movie)
			return
		}
	}

	json.NewEncoder(w).Encode(&Movie{})
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Golang Movie REST")
}

func routers(r *mux.Router) {
	r.HandleFunc("/", home).Methods("GET")
	r.HandleFunc("/api/movies", showMovies).Methods("GET")
	r.HandleFunc("/api/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/api/movies", addMovie).Methods("POST")
}

func main() {
	// Init router
	r := mux.NewRouter();

	// Set routers
	routers(r)

	// Listen port 8080
	log.Fatal(http.ListenAndServe(":8080", r))
}
