package validators

import (
	"api_short_long_polling/src/customers/domain"
	"errors"
)


func CheckCustomer(customer domain.Customer) error {

	if customer.Id < 0 {
		return errors.New("El ID del cliente debe ser mayor o igual a 0")
	}

	if customer.Name == "" {
		return errors.New("El nombre del cliente no puede estar vacío")
	}

	if customer.Last_name == "" {
		return errors.New("El apellido del cliente no puede estar vacío")
	}

	if len(customer.Phone_number) != 10 {
		return errors.New("El número de teléfono debe tener exactamente 10 dígitos")
	}

	if customer.Curp == "" || len(customer.Curp) != 18 {
		return errors.New("El CURP debe tener exactamente 18 caracteres")
	}

	if customer.Number_license == "" {
		return errors.New("El número de licencia no puede estar vacío")
	}

	return nil
}
