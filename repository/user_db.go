package repository

import (
	"gorm.io/gorm"
)

type userRepositoryDB struct {
	db *gorm.DB
}

func NewUserRepositoryDB(db *gorm.DB) userRepositoryDB {
	return userRepositoryDB{
		db: db,
	}
}

func (r userRepositoryDB) GetAll() ([]User, error) {

	var result []User
	r.db.Find(&User{}).Scan(&result)

	return result, nil
}

func (r userRepositoryDB) GetByID(id int) (*User, error) {

	var result User
	r.db.Find(&User{}, "id = ?", id).Scan(&result)

	return &result, nil
}

func (r userRepositoryDB) Create(user User) (*User, error) {

	r.db.Create(&user)

	return &user, nil
}
