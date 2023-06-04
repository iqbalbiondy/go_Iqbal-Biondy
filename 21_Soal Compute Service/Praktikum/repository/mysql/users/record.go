package users

import (
	domain_users "Praktikum/domain/users"
	"gorm.io/gorm"
)

type User struct {
	*gorm.Model
	ID       int
	Email    string
	Password string
}

func ToDomainUsers(rec User) domain_users.Users {
	return domain_users.Users{
		ID:       rec.ID,
		Email:    rec.Email,
		Password: rec.Password,
	}
}
