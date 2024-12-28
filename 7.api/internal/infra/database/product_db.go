package database

import (
	"gorm.io/gorm"

	"github.com/rodrigueslg/fc-goexpert/api/internal/entity"
)

type ProductDB struct {
	DB *gorm.DB
}

func NewProductDB(db *gorm.DB) *ProductDB {
	return &ProductDB{DB: db}
}

func (db *ProductDB) Create(product *entity.Product) error {
	return db.DB.Create(product).Error
}

func (db *ProductDB) FindByID(id string) (*entity.Product, error) {
	var product entity.Product
	err := db.DB.First(&product, "id = ?", id).Error
	return &product, err
}

func (db *ProductDB) FindAll(page, limit int, sort string) ([]entity.Product, error) {
	var products []entity.Product

	if sort != "" && sort != "asc" && sort != "desc" {
		sort = "asc"
	}

	if page != 0 && limit != 0 {
		page := (page - 1) * limit
		err := db.DB.Limit(limit).Offset(page).Order("created_at " + sort).Find(&products).Error
		return products, err
	}

	err := db.DB.Order("created_at " + sort).Find(&products).Error
	return products, err
}

func (db *ProductDB) Update(product *entity.Product) error {
	_, err := db.FindByID(product.ID.String())
	if err != nil {
		return err
	}
	return db.DB.Save(product).Error
}

func (db *ProductDB) Delete(id string) error {
	product, err := db.FindByID(id)
	if err != nil {
		return err
	}
	return db.DB.Delete(product).Error
}
