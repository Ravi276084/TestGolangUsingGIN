package domain

import "github.com/Ravi276084/TestGoLangUsingGIN/GoProgramsUsingGIN/errs"

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, *errs.AppError) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{1001, "Prasad", "Hyd", "500032", "04-AUG-1996", "1"},
		{1002, "Ravi", "Hyd", "500032", "08-AUG-1997", "1"},
	}
	return CustomerRepositoryStub{customers}
}
