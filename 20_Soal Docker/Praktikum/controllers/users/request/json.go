package request

import domain_users "Praktikum/domain/users"

type RequestJSONUsers struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func ToDomainUsers(req RequestJSONUsers) domain_users.Users {
	return domain_users.Users{
		Email:    req.Email,
		Password: req.Password,
	}
}
