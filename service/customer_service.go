package service

import (
	"bank/errs"
	"bank/logs"
	"bank/repository"
)

type customerService struct {
	custRepo repository.CustomerRepository
}

func NewCustomerService(custRepo repository.CustomerRepository) customerService {
	return customerService{custRepo: custRepo}
}

func (s customerService) GetCustomers() ([]CustomerResponse, error) {

	customers, err := s.custRepo.GetAll()
	if err != nil {
		logs.Error(err)
		return nil, errs.NewNotFoundError("customer not found")
	}

	custResponses := []CustomerResponse{}
	for _, customer := range customers {
		custResponse := CustomerResponse{
			CustomerID: customer.CustomerID,
			Name:       customer.Name,
			Status:     customer.Status,
		}
		custResponses = append(custResponses, custResponse)
	}

	return custResponses, nil
}

func (s customerService) GetCustomer(id string) (*CustomerResponse, error) {

	customer, err := s.custRepo.GetByID(id)
	if err != nil {
		logs.Error(err)
		return nil, errs.NewNotFoundError("customer not found")
	}

	CustomerResponse := CustomerResponse{
		CustomerID: customer.CustomerID,
		Name:       customer.Name,
		Status:     customer.Status,
	}

	return &CustomerResponse, nil
}
