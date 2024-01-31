package domain

import "github.com/Ravi276084/TestGoLangUsingGIN/GoProgramsUsingGIN/errs"

type Customer struct {
	Id          int
	Name        string `json:"full_name" xml:"name"`
	City        string `json:"city" xml:"city"`
	Zipcode     string `json:"zipcode" xml:"zipcode"`
	DateofBrith string
	Status      string
}

type CustomerRepository interface {
	FindAll() ([]Customer, *errs.AppError)
	ById(int) (*Customer, *errs.AppError)
}
