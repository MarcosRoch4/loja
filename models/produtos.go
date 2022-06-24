package models

import "loja/db"

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func BuscaProdutos() []Produto {

	db := db.ConectDB()
	defer db.Close()

	produtosSelectAll, err := db.Query("SELECT * FROM produtos")

	if err != nil {
		panic(err.Error())
	}

	p := Produto{}
	produtos := []Produto{}

	for produtosSelectAll.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = produtosSelectAll.Scan(&id, &nome, &descricao, &preco, &quantidade)

		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)

	}
	return produtos
}

func CriarNovoProduto(nome, descricacao string, preco float64, quantidade int) {
	db := db.ConectDB()
	defer db.Close()

	insereDados, err := db.Prepare("INSERT INTO produtos (nome,descricao,preco,quantidade) VALUES ($1,$2,$3,$4)")
	if err != nil {
		panic(err.Error())
	}

	insereDados.Exec(nome, descricacao, preco, quantidade)

}

func DeletaProduto(id string) {
	db := db.ConectDB()
	defer db.Close()

	deletarProduto, err := db.Prepare("DELETE FROM produtos WHERE id = $1")
	if err != nil {
		panic(err.Error())
	}

	deletarProduto.Exec(id)

}
