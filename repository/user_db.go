package repository

import (
	"sync"

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

func (r userRepositoryDB) GetAll(skip, limit int) (Users, error) {

	query := r.db.Raw("SELECT * FROM users LIMIT ? OFFSET ?", limit, skip)
	sum := r.db.Raw("SELECT COUNT(*) FROM users")
	if query.Error != nil || sum.Error != nil {
		return Users{}, query.Error
	}

	var result Users

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		query.Scan(&result.List)
		defer wg.Done()
	}()
	go func() {
		sum.Scan(&result.Total)
		defer wg.Done()
	}()
	wg.Wait()

	return result, nil
}

func (r userRepositoryDB) GetByID(id int) (*User, error) {

	query := r.db.Raw("SELECT * FROM users WHERE id = ? LIMIT 1", id)
	if query.Error != nil {
		return nil, query.Error
	}

	var result User
	query.Scan(&result)

	return &result, nil
}

func (r userRepositoryDB) GetByUser(user string) (*User, error) {

	query := r.db.Raw("SELECT * FROM users WHERE username = ? LIMIT 1", user)
	if query.Error != nil {
		return nil, query.Error
	}

	var result User
	query.Scan(&result)

	return &result, nil
}

func (r userRepositoryDB) Create(user User) (*User, error) {

	if err := r.db.AutoMigrate(&User{}); err != nil {
		return nil, err
	}

	query := r.db.Create(&user)
	if query.Error != nil {
		return nil, query.Error
	}

	return &user, nil
}
