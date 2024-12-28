package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Supplier struct {
	ID       int `gorm:"primaryKey"`
	Name     string
	Products []Product
}

type Category struct {
	ID       int `gorm:"primaryKey"`
	Name     string
	Products []Product `gorm:"many2many:products_categories;"`
}

type Product struct {
	ID           int `gorm:"primaryKey"`
	Name         string
	Price        float64
	SupplierID   int
	Supplier     Supplier
	Categories   []Category `gorm:"many2many:products_categories;"`
	SerialNumber SerialNumber
	gorm.Model
}

type SerialNumber struct {
	ID        int `gorm:"primaryKey"`
	Number    string
	ProductID int
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Product{}, &Supplier{}, &Category{}, &SerialNumber{})

	//create supplier
	s := Supplier{Name: "HP"}
	db.Create(&s)

	//create category 1
	// c1 := Category{Name: "Electronics"}
	// db.Create(&c1)

	//create category 2
	// c2 := Category{Name: "Peripherals"}
	// db.Create(&c2)

	//create product
	// db.Create(&Product{
	// 	Name:       "Notebook",
	// 	Price:      1899.99,
	// 	SupplierID: s.ID,
	// 	Categories: []Category{c1, c2}})

	//create single serial number
	// sn := SerialNumber{
	// 	Number:    "123456",
	// 	ProductID: 1}
	// db.Create(&sn)

	var categories []Category
	err = db.Model(&Category{}).Preload("Products.SerialNumber").Find(&categories).Error
	if err != nil {
		panic(err)
	}
	for _, category := range categories {
		fmt.Println(category.Name, ":")
		for _, product := range category.Products {
			println("-", product.Name, "Serial Number:", product.SerialNumber.Number)
		}
	}

	// create batch
	// db.Create(&[]Product{
	// 	{Name: "Mouse", Price: 99.99, CategoryID: c.ID},
	// 	{Name: "Keyboard", Price: 199.99, CategoryID: c.ID},
	// 	{Name: "Monitor", Price: 799.99, CategoryID: c.ID},
	// })

	// select by id
	// var product Product
	// db.First(&product, 2)
	// fmt.Println(product)

	// select condition
	// var product Product
	// db.First(&product, "name = ?", "Mouse")
	// fmt.Println(product)

	// select where
	// var products []Product
	// db.Where("price > ?", 1000).Find(&products)
	// for _, product := range products {
	// 	fmt.Println(product)
	// }

	// select all
	// var products []Product
	// db.Preload("Category").Preload("SerialNumber").Find(&products)
	// for _, product := range products {
	// 	fmt.Println(product.Name, product.Category.Name, product.SerialNumber.Number)
	// }

	// select paginated
	// var products []Product
	// db.Limit(2).Offset(2).Find(&products)
	// for _, product := range products {
	// 	fmt.Println(product)
	// }

	// update
	// var p Product
	// db.First(&p, 1)
	// p.Name = "New Notebook"
	// db.Save(&p)
	// var p2 Product
	// db.First(&p2, 1)
	// fmt.Println(p2)

	// delete
	// var p Product
	// db.First(&p, 1)
	// db.Delete(&p)
}
