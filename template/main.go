package main

import (
	"Main_project/greeter_client"
	"fmt"
	"html/template"
	"net/http"
)

func create(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./create.html", "./header.html", "./footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	tmpl.ExecuteTemplate(w, "create", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./index.html", "./header.html", "./footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	tmpl.ExecuteTemplate(w, "index", nil)
}

func add(x, y int) int {
	return x + y
}

func saveArticle(w http.ResponseWriter, r *http.Request) {
	//Adding a function name to FuncMap for searching
	funcs := template.FuncMap{"add": add}
	//Parses files then adds the elements of argument FuncMap to the template's function map , allocates new HTML template and wraps a call to a function returning
	tmpl := template.Must(template.New("index1.html").Funcs(funcs).ParseFiles("./save_article.html", "./header.html", "./footer.html"))
	//returns the first value of the name component of the query
	title := r.FormValue("title")
	//are calling a function Do_deal() for connect to server, send a youtube API and retrieving a data from server
	Data := greeter_client.Do_deal(title)

	if title == "" {
		fmt.Fprintf(w, "What gonna being wrong!")
	}
	tmpl.ExecuteTemplate(w, "save_article", Data)
}

func handleFunc() {
	http.HandleFunc("/", index)
	http.HandleFunc("/create", create)
	http.HandleFunc("/save_article", saveArticle)
	http.ListenAndServe("0.0.0.0:8080", nil)
}

func main() {
	handleFunc()
}
