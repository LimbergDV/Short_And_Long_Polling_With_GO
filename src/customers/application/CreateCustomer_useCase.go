package application

import "api_short_long_polling/src/customers/domain"



type CreateCustomer struct{
	db domain.ICustomer
}

func NewCreateCustomer (db domain.ICustomer) *CreateCustomer {
	return &CreateCustomer{db: db}
}

func (cc *CreateCustomer) Run (customer domain.Customer) (uint, error) {
	return cc.db.Save(customer)
}