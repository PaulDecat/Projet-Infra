package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func indexhandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/index.html")
}

func submitPostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		content := r.FormValue("content")

		_, err := db.Exec("INSERT INTO Save (Content) VALUES (?)", content)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/index", http.StatusSeeOther)
	}
}

func main() {
	var err error
	db, err = initDB("./save.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	fmt.Printf("Starting server at http://localhost:8080")
	http.HandleFunc("/", indexhandler)
	http.HandleFunc("/submitpost", submitPostHandler)
	http.ListenAndServe(":8080", nil)

}
