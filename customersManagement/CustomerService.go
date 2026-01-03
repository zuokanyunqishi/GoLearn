package main

type CustomerService struct {
	customers   []CustomerModel
	customerNum uint
}

func (s *CustomerService) Add(customer CustomerModel) bool {
	s.customerNum++
	customer.Id = s.customerNum
	s.customers = append(s.customers, customer)
	return true

}

func (s *CustomerService) List() []CustomerModel {
	return s.customers
}

func (s *CustomerService) Update(index int, model CustomerModel) {
	s.customers[index].Phone = model.Phone
	s.customers[index].Name = model.Name
	s.customers[index].Age = model.Age
	s.customers[index].Gender = model.Gender
	s.customers[index].Email = model.Email

}

func (s *CustomerService) DeleteCustomer(index int) {
	s.customers = append(s.customers[:index], s.customers[index+1:]...)
	return
}

func NewCustomerService() *CustomerService {
	return &CustomerService{
		customers: []CustomerModel{},
	}
}
