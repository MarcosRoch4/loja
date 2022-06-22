package models


import "github.com/MarcosRoch4/db"

type Produto struct{
	id int64
	Nome string
	Descricao string
	Preco float64
	Quantidade int64
}

func BuscaProdutos(){

	db := db.ConectDB()
	defer db.Close()

	produtosSelectAll, err := db.Query("SELECT * FROM produtos")

	if err != nil {
		panic(err.Error())
	}
	
	p :=  Produto{}
	produtos := []Produto{}

	for produtosSelectAll.Next(){
		var id, quantidade int64
		var nome,descricao string
		var preco float64

		err = produtosSelectAll.Scan(&id,&nome,&descricao,&preco,&quantidade)

		if err != nil {
			panic(err.Error())
		}

		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos,p)

	}
	return produtos
}