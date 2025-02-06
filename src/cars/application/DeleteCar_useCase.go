package application

import "api_short_long_polling/src/cars/domain"

type DeleteCar struct{
	db domain.ICar
}

func NewDeleteCar (db domain.ICar) *DeleteCar {
	return &DeleteCar{db: db}
}

func (dc *DeleteCar) Run (id int) (uint, error) {
	return dc.db.Delete(id)
}