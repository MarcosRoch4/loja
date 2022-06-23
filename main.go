package main

import (
	"html/template"
	"net/http"
	"github.com/MarcosRoch4/loja/models"
	"github.com/MarcosRoch4/loja/db"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {	
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)

}

func index(w http.ResponseWriter, r *http.Request) {
	
	allProdutos := models.BuscaProdutos()
	

	temp.ExecuteTemplate(w, "Index", allProdutos)
}

