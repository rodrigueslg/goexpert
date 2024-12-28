package database

import (
	"gorm.io/gorm"

	"github.com/rodrigueslg/fc-goexpert/api/internal/entity"
)

type UserDB struct {
	DB *gorm.DB
}

func NewUserDB(db *gorm.DB) *UserDB {
	return &UserDB{
		DB: db,
	}
}

func (db *UserDB) Create(user *entity.User) error {
	return db.DB.Create(user).Error
}

func (db *UserDB) FindByEmail(email string) (*entity.User, error) {
	var user entity.User
	if err := db.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
