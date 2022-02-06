package models

import (
	"github.com/marcos-nsantos/loja-go/database"
	"log"
)

type Produto struct {
	ID         uint64
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade uint64
}

func BuscarTodosOsProdutos() []Produto {
	db, err := database.Conectar()
	if err != nil {
		log.Panic(err.Error())
	}
	defer db.Close()

	todosOsProdutosDoBancoDeDados, err := db.Query("SELECT * FROM produto ORDER BY id ASC")
	if err != nil {
		log.Panic(err.Error())
	}
	defer todosOsProdutosDoBancoDeDados.Close()

	var produtos []Produto
	for todosOsProdutosDoBancoDeDados.Next() {
		produto := Produto{}

		err := todosOsProdutosDoBancoDeDados.Scan(&produto.ID, &produto.Nome, &produto.Descricao, &produto.Preco, &produto.Quantidade)
		if err != nil {
			log.Panic(err.Error())
		}

		produtos = append(produtos, produto)
	}
	return produtos
}
