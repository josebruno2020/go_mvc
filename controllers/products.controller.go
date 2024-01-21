package controllers

import (
	"log"
	"net/http"
	"strconv"
	"text/template"
	"webapp/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	products := models.GetAllProducts()

	temp.ExecuteTemplate(w, "index", products)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "new", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", 301)
	}
	name := r.FormValue("nome")
	description := r.FormValue("descricao")
	price := r.FormValue("preco")
	amount := r.FormValue("quantidade")

	priceFloat, err := strconv.ParseFloat(price, 64)

	if err != nil {
		log.Printf("Erro na conversão do preço: %s", err.Error())
	}

	amountInt, err := strconv.Atoi(amount)

	if err != nil {
		log.Printf("Erro na conversão da quantidade: %s", err.Error())
	}

	models.SaveProduct(name, description, priceFloat, amountInt)

	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		http.Redirect(w, r, "/", 301)
	}
	productId := r.URL.Query().Get("id")
	models.DeleteProduct(productId)
	http.Redirect(w, r, "/", 301)
}
