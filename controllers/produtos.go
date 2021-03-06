package controllers

import (
	"github.com/marcos-nsantos/loja-go/models"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

var templates = template.Must(template.ParseGlob("templates/*gohtml"))

func Index(w http.ResponseWriter, r *http.Request) {
	todosOsProdutos := models.BuscarTodosOsProdutos()
	err := templates.ExecuteTemplate(w, "Index", todosOsProdutos)
	if err != nil {
		log.Println(err)
	}
}

func New(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "New", nil)
	if err != nil {
		log.Println(err)
	}
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		precoStringParaFloat64, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Println("Erro na conversão de preco:", err)
		}

		quantidadeStringParaUint64, err := strconv.ParseUint(quantidade, 10, 64)
		if err != nil {
			log.Println("Erro na conversão de quantidade:", err)
		}

		models.CriarNovoProduto(nome, descricao, precoStringParaFloat64, quantidadeStringParaUint64)
	}

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	produtoID := r.URL.Query().Get("id")
	models.DeletarProduto(produtoID)
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	produtoID := r.URL.Query().Get("id")
	produto := models.DetalhesDoProduto(produtoID)
	err := templates.ExecuteTemplate(w, "Edit", produto)
	if err != nil {
		log.Println(err)
	}
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		ID := r.FormValue("id")
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		IDStringParaUint64, err := strconv.ParseUint(ID, 10, 64)
		if err != nil {
			log.Panic("Erro no conversao de ID:", err)
		}

		precoStringParaFloat64, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Println("Erro na conversão de preco:", err)
		}

		quantidadeStringParaUint64, err := strconv.ParseUint(quantidade, 10, 64)
		if err != nil {
			log.Println("Erro na conversão de quantidade:", err)
		}

		models.AtualizarProduto(IDStringParaUint64, nome, descricao, precoStringParaFloat64, quantidadeStringParaUint64)
	}

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}
