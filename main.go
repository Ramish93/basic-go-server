package main

import (
	"fmt"
	"log"

	// "encoding/json"
	// "math/rand"
	"net/http"
	// "strconv"
	"github.com/gorilla/mux"
)

//____________________code for server to serve static folder ___________

// func formHandler(w http.ResponseWriter, r *http.Request){
// 	if err := r.ParseForm(); err != nil {
// 		fmt.Fprintf(w, "parse form error: %v", err)
// 		return
// 	}
// 	fmt.Fprintf(w, "Post req successfully ")
// 	name := r.FormValue("name")
// 	address := r.FormValue("address")
// 	fmt.Fprintf(w, "name = %s ", name)
// 	fmt.Fprintf(w, "address = %s ", address)

// }

// func helloHandler(w http.ResponseWriter, r *http.Request) {
// 	if r.URL.Path != "/hello" {
// 		http.Error(w, "404 Not Found", http.StatusNotFound)
// 		return
// 	}
// 	if r.Method != "GET" {
// 		http.Error(w, "method not supported", http.StatusNotFound)
// 		return
// 	}

// 	fmt.Fprintf(w, "Hello Go")
// }

// __________________________code for CRUD movie API below____________________

type Movie struct {
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Director *Director `json:"director"`
}
type Director struct{
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
}

var movies []Movie

func main () {
	// __________________________code for CRUD movie API below____________________

	r := mux.NewRouter()

	movies = append(movies, Movie{
		ID: "1",
		Isbn:"438227",
		Title: "Movie one",
		Director: &Director{
			FirstName: "James",
			LastName: "Camron",
		},

	})
	movies = append(movies, Movie{
		ID: "2",
		Isbn: "438228",
		Title: "Movie two",
		Director: &Director{
			FirstName: "James",
			LastName: "Di Caprio",
		},
	})
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies/{id}", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovies).Methods("PUT")
	r.HandleFunc("movies/{id}", deleteMovies).Methods("Delete")

	fmt.Printf("staring server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))


//____________________code for server to serve static folder ___________
	// fileServer := http.FileServer(http.Dir("./static"))
	// http.Handle("/", fileServer)
	// http.HandleFunc("/form", formHandler)
	// http.HandleFunc("/hello", helloHandler)

	// fmt.Printf("strting server at port 8080\n")
	// if err:= http.ListenAndServe(":8080", nil); err != nil {
	// 	log.Fatal(err)
	// }

	
}