package application

import "api_short_long_polling/src/customers/domain"



type DeleteCustomer struct{
	db domain.ICustomer
}

func NewDeleteCustomer (db domain.ICustomer) *DeleteCustomer {
	return &DeleteCustomer{db: db}
}

func (dc *DeleteCustomer) Run (id int) (uint, error) {
	return dc.db.Delete(id)
}