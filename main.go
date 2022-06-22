package main

import (
	"html/template"
	"net/http"
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/MarcosRoch4/loja/models"
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

