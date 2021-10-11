package user

import "gorm.io/gorm"

type Repository interface {
	Store(user User) (User, error)
	FindByEmail(email string) (User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) Store(user User) (User, error) {
	err := r.db.Create(&user).Error
	return user, err
}

func (r *repository) FindByEmail(email string) (User, error) {
	var user User
	err := r.db.Where("email = ?", email).Find(&user).Error
	return user, err
}


