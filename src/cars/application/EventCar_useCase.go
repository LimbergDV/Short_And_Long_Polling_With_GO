package application

import "api_short_long_polling/src/cars/domain"

type GetAvailableCars struct {
    db domain.ICar
}

func NewGetAvailableCars(db domain.ICar) *GetAvailableCars {
    return &GetAvailableCars{db: db}
}

func (u *GetAvailableCars) Run() []domain.Car {
    allCars := u.db.GetAll()
    var availableCars []domain.Car
    for _, car := range allCars {
        if car.Available { // asumiendo que la entidad Car tiene un campo Available bool
            availableCars = append(availableCars, car)
        }
    }
    return availableCars
}
