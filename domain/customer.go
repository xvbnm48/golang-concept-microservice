package domain

import "github.com/xvbnm48/golang-concept-microservice/errs"

type Customer struct {
	Id          string
	Name        string
	City        string
	ZipCode     string
	DateOfBirth string
	Status      string
}

type CustomerRepository interface {
	FindAll() ([]Customer, error)
	ById(id string) (*Customer, *errs.AppError)
}
