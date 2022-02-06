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

func CriarNovoProduto(nome string, descricao string, preco float64, quantidade uint64) {
	db, err := database.Conectar()
	if err != nil {
		log.Panic(err.Error())
	}
	defer db.Close()

	insereDadosNoBanco, err := db.Prepare("INSERT INTO produto(nome, descricao, preco, quantidade) VALUES ($1, $2, $3, $4)")
	if err != nil {
		log.Panic(err.Error())
	}
	defer insereDadosNoBanco.Close()

	_, err = insereDadosNoBanco.Exec(nome, descricao, preco, quantidade)
	if err != nil {
		log.Panic(err.Error())
	}
}

func DeletarProduto(ID string) {
	db, err := database.Conectar()
	if err != nil {
		log.Panic(err.Error())
	}
	defer db.Close()

	deletaProduto, err := db.Prepare("DELETE FROM produto WHERE id=$1")
	if err != nil {
		log.Panic(err)
	}
	defer deletaProduto.Close()

	_, err = deletaProduto.Exec(ID)
	if err != nil {
		log.Panic(err.Error())
	}
}

func DetalhesDoProduto(ID string) Produto {
	db, err := database.Conectar()
	if err != nil {
		log.Panic(err.Error())
	}
	defer db.Close()

	produtoDoBanco, err := db.Query("SELECT * FROM produto WHERE id=$1", ID)
	if err != nil {
		log.Panic(err.Error())
	}
	defer produtoDoBanco.Close()

	produto := Produto{}
	for produtoDoBanco.Next() {
		err := produtoDoBanco.Scan(&produto.ID, &produto.Nome, &produto.Descricao, &produto.Preco, &produto.Quantidade)
		if err != nil {
			log.Panic(err.Error())
		}
	}
	return produto
}

func AtualizarProduto(ID uint64, nome string, descricao string, preco float64, quantidade uint64) {
	db, err := database.Conectar()
	if err != nil {
		log.Panic(err.Error())
	}
	defer db.Close()

	atualizaProduto, err := db.Prepare("UPDATE produto SET nome=$1, descricao=$2, preco=$3, quantidade=$4 WHERE id=$5")
	if err != nil {
		log.Panic(err.Error())
	}
	defer atualizaProduto.Close()

	_, err = atualizaProduto.Exec(nome, descricao, preco, quantidade, ID)
	if err != nil {
		log.Panic(err.Error())
	}
}
