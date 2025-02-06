package application

import "api_short_long_polling/src/cars/domain"



type UpdateCar struct {
	db domain.ICar
}

func NewUpdateCar( db domain.ICar) *UpdateCar {
	return &UpdateCar{db: db}
}

func (ue *UpdateCar) Run (id int, car domain.Car) (uint, error) {
	return ue.db.Update(id, car)
}