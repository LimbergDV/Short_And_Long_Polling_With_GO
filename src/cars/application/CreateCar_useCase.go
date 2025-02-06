package application

import "api_short_long_polling/src/cars/domain"

type CreateCar struct{
	db domain.ICar
}

func NewCreateCar (db domain.ICar) *CreateCar {
	return &CreateCar{db: db}
}

func (cc *CreateCar) Run (car domain.Car) (uint, error) {
	return cc.db.Save(car)
}