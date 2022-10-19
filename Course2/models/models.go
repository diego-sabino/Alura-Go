package models

import (
	"todolist-go/db"
)

type Product struct {
	Id          int
	Name        string
	Description string
	Price       int
	Quantity    int
}

func FindAllProducts() []Product {
	db := db.ConnectSQL()

	selectAllProducts, err := db.Query("select * from products order by id asc")
	if err != nil {
		panic(err.Error())
	}

	p := Product{}
	Products := []Product{}

	for selectAllProducts.Next() {
		var Id, Quantity, Price int
		var Name, Description string

		err = selectAllProducts.Scan(&Id, &Name, &Description, &Price, &Quantity)
		if err != nil {
			panic(err.Error())
		}

		p.Id = Id
		p.Name = Name
		p.Description = Description
		p.Quantity = Quantity
		p.Price = Price

		Products = append(Products, p)
	}
	defer db.Close()
	return Products
}

func CreateNewProduct(name, description string, price int, quantity int) {
	db := db.ConnectSQL()

	insertData, err := db.Prepare("insert into products (name, description, price, quantity) values($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}

	insertData.Exec(name, description, price, quantity)
	defer db.Close()
}

func DeleteProduct(id string) {
	db := db.ConnectSQL()

	deleteData, err := db.Prepare("delete from products where id=$1")
	if err != nil {
		panic(err.Error())
	}

	deleteData.Exec(id)
	defer db.Close()
}

func EditProduct(id string) Product {
	db := db.ConnectSQL()
	selData, err := db.Query("select * from products where id=$1", id)
	if err != nil {
		panic(err.Error())
	}

	editProduct := Product{}

	for selData.Next() {
		var Id, Quantity, Price int
		var Name, Description string

		err = selData.Scan(&Id, &Name, &Description, &Price, &Quantity)
		if err != nil {
			panic(err.Error())
		}

		editProduct.Id = Id
		editProduct.Name = Name
		editProduct.Description = Description
		editProduct.Quantity = Quantity
		editProduct.Price = Price
	}
	defer db.Close()
	return editProduct
}

func UpdateProduct(id int, name, description string, price, quantity int) {
	db := db.ConnectSQL()

	updateData, err := db.Prepare("update products set name=$1, description=$2, price=$3, quantity=$4 where id=$5")
	if err != nil {
		panic(err.Error())
	}
	updateData.Exec(name, description, price, quantity, id)
	defer db.Close()
}
