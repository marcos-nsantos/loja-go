package controllers

import (
	"github.com/marcos-nsantos/loja-go/models"
	"html/template"
	"net/http"
)

var templates = template.Must(template.ParseGlob("templates/*gohtml"))

func Index(w http.ResponseWriter, r *http.Request) {
	todosOsProdutos := models.BuscarTodosOsProdutos()
	templates.ExecuteTemplate(w, "Index", todosOsProdutos)
}
