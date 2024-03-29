package service

import (
	"github.com/Ravi276084/TestGoLangUsingGIN/GoProgramsUsingGIN/domain"
	"github.com/Ravi276084/TestGoLangUsingGIN/GoProgramsUsingGIN/errs"
)

//go:generate mockgen -destination= C:/Users/PANRAO/GoPrograms/mocks/service/mockCustomerService.go -package=service github.com/ashishjuyal/banking/service CustomerService
type CustomerService interface {
	GetAllCustomers() ([]domain.Customer, *errs.AppError)
	GetCustomer(int) (*domain.Customer, *errs.AppError)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomers() ([]domain.Customer, *errs.AppError) {
	return s.repo.FindAll()
}

func (s DefaultCustomerService) GetCustomer(id int) (*domain.Customer, *errs.AppError) {
	return s.repo.ById(id)
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}
}
