package models

import (
	"database/sql"
	"webapp/db"
)

type Product struct {
	Id                int
	Name, Description string
	Price             float64
	Amount            int
}

func GetAllProducts() []Product {
	db := db.ConnectDB()

	selectProducts, err := db.Query("select * from produtos order by id asc")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	return scanSqlProduct(selectProducts)
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

func GetProductById(id string) Product {
	db := db.ConnectDB()

	sqlQuery, err := db.Query("select * from produtos where id = $1", id)

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	products := scanSqlProduct(sqlQuery)

	return products[0]
}

func UpdateProductById(product Product) {
	db := db.ConnectDB()

	sqlUpdate, err := db.Prepare("update produtos set nome = $1, descricao = $2, preco = $3, quantidade = $4 where id = $5")

	if err != nil {
		panic(err.Error())
	}

	sqlUpdate.Exec(product.Name, product.Description, product.Price, product.Amount, product.Id)

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

func scanSqlProduct(rows *sql.Rows) []Product {
	products := []Product{}
	p := Product{}

	for rows.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err := rows.Scan(&id, &nome, &descricao, &preco, &quantidade)

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

	return products
}
