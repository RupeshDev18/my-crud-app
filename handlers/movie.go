package handlers

import (
	"encoding/json"
	"github.com/rupeshdev18/my-crud-app/models"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var movies []models.Movie

func InitializeMovies() {
	movies = append(movies, models.Movie{ID: "1", Isbn: "892374", Title: "MovieNo1", Director: &models.Director{Firstname: "John", Lastname: "Doe"}})
	movies = append(movies, models.Movie{ID: "2", Isbn: "049374", Title: "MovieNo2", Director: &models.Director{Firstname: "Steven", Lastname: "Smith"}})
}

func GetMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func GetMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie models.Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(100000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func UpdateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var movie models.Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = strconv.Itoa(rand.Intn(100000))
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
		}
	}
}
