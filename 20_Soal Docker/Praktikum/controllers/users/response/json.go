package response

import domain_users "Praktikum/domain/users"

type ResponseJSONUsers struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func FromDomainUsers(domain domain_users.Users) ResponseJSONUsers {
	return ResponseJSONUsers{
		Email:    domain.Email,
		Password: domain.Password,
	}
}
