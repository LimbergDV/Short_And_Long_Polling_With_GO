package infrastructure

import (
	core "api_short_long_polling/src/Core"
	"api_short_long_polling/src/cars/domain"
	"fmt"
	"log"
)

// alamacena la conexión a la bd
type MySQL struct {
	conn *core.Conn_MySQL
}

func NewMySQL() *MySQL {
	conn := core.GetDBPool()
	if conn.Err != "" {
		log.Fatalf("Error al configurar el pool de conexiones: %v", conn.Err)
	}
	return &MySQL{conn: conn}
}

func (mysql *MySQL) Save(car domain.Car) (uint, error) {
	query := "INSERT INTO cars (brand, model, year, type_car, plate_number, price_day) VALUES (?, ?, ?, ?, ?, ?)"
	res, err := mysql.conn.ExecutePreparedQuery(query, car.Brand, car.Model, car.Year, car.Type_Car, car.Plate_number, car.Price_day)
	if err != nil {
		fmt.Println("Error al preparar la consulta:", err)
		return 0, err
	}
	id, _ := res.LastInsertId()
	fmt.Println("Carro creado")
	return uint(id), nil
}

func (mysql *MySQL) GetAll() []domain.Car {
	query := "SELECT * FROM cars"
	var cars []domain.Car
	rows := mysql.conn.FetchRows(query)

	if rows == nil {
		fmt.Println("No se obtuvieron los datos")
		return cars
	}
	defer rows.Close()

	for rows.Next() {
		var c domain.Car
		if err := rows.Scan(&c.Id, &c.Brand, &c.Model, &c.Year, &c.Type_Car, &c.Plate_number, &c.Price_day); err != nil {
			fmt.Println("Error al escanear la fila:", err)
		} else {
			cars = append(cars, c)
		}
	}

	if err := rows.Err(); err != nil {
		fmt.Println("Error iterando sobre las filas:", err)
	}

	fmt.Println("Lista de carros obtenida")
	return cars
}

func (mysql *MySQL) Delete(id int) (uint, error) {
	query := "DELETE FROM cars WHERE id = ?"
	res, err := mysql.conn.ExecutePreparedQuery(query, id)
	if err != nil {
		fmt.Println("Error al ejecutar la consulta:", err)
		return 0, err
	}
	rows, _ := res.RowsAffected()
	if rows == 0 {
		return 0, fmt.Errorf("No se encontró ningún carro con el ID proporcionado")
	}
	fmt.Println("Carro eliminado")
	return uint(rows), nil
}

func (mysql *MySQL) Update(id int, car domain.Car) (uint, error) {
	query := "UPDATE cars SET brand = ?, model = ?, year = ?, type_car = ?, plate_number = ?, price_day = ? WHERE id = ?"
	res, err := mysql.conn.ExecutePreparedQuery(query, car.Brand, car.Model, car.Year, car.Type_Car, car.Plate_number, car.Price_day, id)
	if err != nil {
		fmt.Println("Error al ejecutar la consulta:", err)
		return 0, err
	}
	rows, _ := res.RowsAffected()
	if rows == 0 {
		return 0, fmt.Errorf("No se encontró ningún carro con el ID proporcionado para actualizar")
	}
	fmt.Println("Carro actualizado")
	return uint(rows), nil
}

func (mysql *MySQL) GetById(id int) (domain.Car, error) {
	var car domain.Car

	query := "SELECT id, brand, model, year, type_car, plate_number, price_day FROM cars WHERE id = ?"
	rows := mysql.conn.FetchRows(query, id)

	if rows == nil {
		return car, fmt.Errorf("No se encontraron datos")
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&car.Id, &car.Brand, &car.Model, &car.Year, &car.Type_Car, &car.Plate_number, &car.Price_day)
		if err != nil {
			return car, err
		}
		return car, nil
	}

	return car, fmt.Errorf("No se encontró ningún carro con el ID proporcionado")
}