package main

import (
	"Main_project/greeter_client"
	"fmt"
	"html/template"
	"net/http"
)

func create(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./create.html", "./header.html", "footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	tmpl.ExecuteTemplate(w, "create", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./index.html", "./header.html", "footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	tmpl.ExecuteTemplate(w, "index", nil)
}

func save_article(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	greeter_client.Do_deal(title)

	if title == "" {
		fmt.Fprintf(w, "What gonna being wrong!")
	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func handleFunc() {
	http.HandleFunc("/", index)
	http.HandleFunc("/create", create)
	http.HandleFunc("/save_article", save_article)
	http.ListenAndServe("localhost:8080", nil)
}

func main() {
	handleFunc()

}
