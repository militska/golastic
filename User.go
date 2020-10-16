package main

import "syreclabs.com/go/faker"

type User struct {
	Name     string
	Phone    string
	Address  string
	Company  string
	Birthday string
}

func generateUser() User {
	user := User{
		Address:  faker.Address().City(),
		Name:     faker.Name().Name(),
		Birthday: faker.Date().Birthday(13, 70).String(),
		Phone:    faker.PhoneNumber().CellPhone(),
		Company:  faker.Company().Name(),
	}

	return user
}
