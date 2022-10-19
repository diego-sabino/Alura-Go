package controllers

import (
	"log"
	"net/http"
	"strconv"
	"text/template"
	"todolist-go/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	allProducts := models.FindAllProducts()
	temp.ExecuteTemplate(w, "index", allProducts)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("nome")
		description := r.FormValue("descricao")
		price := r.FormValue("preco")
		quantity := r.FormValue("quantidade")

		convertedPrice, err := strconv.Atoi(price)

		if err != nil {
			log.Println("Erro na conversão do preço:", err)
		}

		convertedQuantity, err := strconv.Atoi(quantity)

		if err != nil {
			log.Println("Erro na conversão da quantidade:", err)
		}

		models.CreateNewProduct(name, description, convertedPrice, convertedQuantity)
	}
	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idProduct := r.URL.Query().Get("id")
	models.DeleteProduct(idProduct)
	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	idProduct := r.URL.Query().Get("id")
	product := models.EditProduct(idProduct)
	temp.ExecuteTemplate(w, "Edit", product)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		name := r.FormValue("nome")
		description := r.FormValue("descricao")
		price := r.FormValue("preco")
		quantity := r.FormValue("quantidade")

		convertedPrice, err := strconv.Atoi(price)
		if err != nil {
			log.Println("Erro na conversão do preço:", err)
		}

		convertedId, err := strconv.Atoi(id)
		if err != nil {
			log.Println("Erro na conversão do ID para int:", err)
		}

		convertedQuantity, err := strconv.Atoi(quantity)
		if err != nil {
			log.Println("Erro na conversão da quantidade:", err)
		}

		models.UpdateProduct(convertedId, name, description, convertedPrice, convertedQuantity)
	}
	http.Redirect(w, r, "/", 301)
}
