package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomersRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{"1", "John", "New York", "10001", "01/01/1990", "active"},
		{"2", "Jane", "New York", "10001", "01/01/1990", "active"},
		{"3", "Jack", "New York", "10001", "01/01/1990", "active"},
	}
	return CustomerRepositoryStub{customers}
}
