package models

import "webapp/db"

type Product struct {
	Id                int
	Name, Description string
	Price             float64
	Amount            int
}

func GetAllProducts() []Product {
	db := db.ConnectDB()

	selectProducts, err := db.Query("select * from produtos")

	if err != nil {
		panic(err.Error())
	}

	p := Product{}
	products := []Product{}

	for selectProducts.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err := selectProducts.Scan(&id, &nome, &descricao, &preco, &quantidade)

		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Name = nome
		p.Description = descricao
		p.Price = preco
		p.Amount = quantidade

		products = append(products, p)
	}
	defer db.Close()

	return products
}

func SaveProduct(name, description string, price float64, amount int) {
	db := db.ConnectDB()

	sqlInsert, err := db.Prepare("insert into produtos(nome, descricao, preco, quantidade) values($1, $2, $3, $4)")

	if err != nil {
		panic(err.Error())
	}

	sqlInsert.Exec(name, description, price, amount)

	defer db.Close()
}

func DeleteProduct(id string) {
	db := db.ConnectDB()

	sqlDelete, err := db.Prepare("delete from produtos where id = $1")

	if err != nil {
		panic(err.Error())
	}

	sqlDelete.Exec(id)

	defer db.Close()
}
