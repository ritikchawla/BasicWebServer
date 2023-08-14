package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Film struct {
	Title    string
	Director string
}

func main() {
	fmt.Println("Go app...")

	http.HandleFunc("/", handleFilmList)
	http.HandleFunc("/add-film/", handleAddFilm)

	log.Fatal(http.ListenAndServe(":8000", nil))
}

func handleFilmList(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("index.html"))
	films := map[string][]Film{
		"Films": {
			{Title: "The Story Of My Experiments With The Truth", Director: "Mahatma Gandhi"},
			{Title: "The Guide", Director: "R.K. Narayan"},
			{Title: "God of Small Things", Director: "Arundhati Roy"},
		},
	}
	tmpl.Execute(w, films)
}

func handleAddFilm(w http.ResponseWriter, r *http.Request) {
	// time.Sleep(1 * time.Second)
	title := r.PostFormValue("title")
	director := r.PostFormValue("director")
	tmpl := template.Must(template.ParseFiles("index.html"))
	tmpl.ExecuteTemplate(w, "film-list-element", Film{Title: title, Director: director})
}
