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
	tmpl, err := template.ParseFiles("./index.html", "./header.html", "./footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	tmpl.ExecuteTemplate(w, "index", nil)
}

func saveArticle(w http.ResponseWriter, r *http.Request) {
	//const templ = `{{range .Items}}----------------------------
	//ID: {{ .Id}}`
	//t := template.Must(template.New("").Parse(templ))
	tmpl, err := template.ParseFiles("./save_article.html", "./header.html", "./footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	title := r.FormValue("title")
	Data := greeter_client.Do_deal(title)
	//fmt.Fprintf(w, data.Kind)
	//fmt.Println(Data)
	if title == "" {
		fmt.Fprintf(w, "What gonna being wrong!")
	}
	tmpl.ExecuteTemplate(w, "save_article", Data)
}

func handleFunc() {
	http.HandleFunc("/", index)
	http.HandleFunc("/create", create)
	http.HandleFunc("/save_article", saveArticle)
	http.ListenAndServe("localhost:8080", nil)
}

func main() {
	handleFunc()
}
