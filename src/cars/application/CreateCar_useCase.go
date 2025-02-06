package application

import "api_short_long_polling/src/cars/domain"

type CreateCar struct{
	db domain.ICar
}

func NewCreateEmployee (db domain.ICar) *CreateCar {
	return &CreateCar{db: db}
}

func (cc *CreateCar) Run (employee domain.Car) (uint, error) {
	return cc.db.Save(employee)
}