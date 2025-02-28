package application

import "api_short_long_polling/src/customers/domain"

// GetAllCustomers es el caso de uso para obtener la lista de todos los customers.
type GetAllCustomersWithChanges struct {
	db domain.ICustomer
}

// NewGetAllCustomers crea una nueva instancia del caso de uso.
func NewGetAllCustomersWitjChanges(db domain.ICustomer) *GetAllCustomers {
	return &GetAllCustomers{db: db}
}


func (uc *GetAllCustomers) Execute() []domain.Customer {
	allCustomers := uc.db.GetAll()
	// Inicializamos el slice para evitar que se serialice como null
	customers := make([]domain.Customer, 0)
	for _, customer := range allCustomers {
		customers = append(customers, customer)
	}
	return customers
}
