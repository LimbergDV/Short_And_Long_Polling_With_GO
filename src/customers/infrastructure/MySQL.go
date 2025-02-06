package infrastructure

import (
	core "api_short_long_polling/src/Core"
	"api_short_long_polling/src/customers/domain"
	"fmt"
	"log"
)

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

func (mysql *MySQL) Save(customer domain.Customer) (uint, error) {
	query := "INSERT INTO customers (name, last_name, phone_number, curp, number_license) VALUES (?, ?, ?, ?, ?)"
	res, err := mysql.conn.ExecutePreparedQuery(query, customer.Name, customer.Last_name, customer.Phone_number, customer.Curp, customer.Number_license)
	if err != nil {
		fmt.Println("Error al preparar la consulta:", err)
		return 0, err
	}
	id, _ := res.LastInsertId()
	fmt.Println("Cliente creado")
	return uint(id), nil
}

func (mysql *MySQL) GetAll() []domain.Customer {
	query := "SELECT * FROM customers"
	var customers []domain.Customer
	rows := mysql.conn.FetchRows(query)

	if rows == nil {
		fmt.Println("No se obtuvieron los datos")
		return customers
	}
	defer rows.Close()

	for rows.Next() {
		var c domain.Customer
		if err := rows.Scan(&c.Id, &c.Name, &c.Last_name, &c.Phone_number, &c.Curp, &c.Number_license); err != nil {
			fmt.Println("Error al escanear la fila:", err)
		} else {
			customers = append(customers, c)
		}
	}

	if err := rows.Err(); err != nil {
		fmt.Println("Error iterando sobre las filas:", err)
	}

	fmt.Println("Lista de clientes obtenida")
	return customers
}

func (mysql *MySQL) Delete(id int) (uint, error) {
	query := "DELETE FROM customers WHERE id = ?"
	res, err := mysql.conn.ExecutePreparedQuery(query, id)
	if err != nil {
		fmt.Println("Error al ejecutar la consulta:", err)
		return 0, err
	}
	rows, _ := res.RowsAffected()
	if rows == 0 {
		return 0, fmt.Errorf("No se encontró ningún cliente con el ID proporcionado")
	}
	fmt.Println("Cliente eliminado")
	return uint(rows), nil
}

func (mysql *MySQL) Update(id int, customer domain.Customer) (uint, error) {
	query := "UPDATE customers SET name = ?, last_name = ?, phone_number = ?, curp = ?, number_license = ? WHERE id = ?"
	res, err := mysql.conn.ExecutePreparedQuery(query, &customer.Name, &customer.Last_name, &customer.Phone_number, &customer.Curp, &customer.Number_license, id)
	if err != nil {
		fmt.Println("Error al ejecutar la consulta:", err)
		return 0, err
	}
	rows, _ := res.RowsAffected()
	if rows == 0 {
		return 0, fmt.Errorf("No se encontró ningún cliente con el ID proporcionado para actualizar")
	}
	fmt.Println("Cliente actualizado")
	return uint(rows), nil
}

func (mysql *MySQL) GetById(id int) (domain.Customer, error) {
	var customer domain.Customer

	query := "SELECT id, name, last_name, phone_number, curp, number_license FROM customers WHERE id = ?"
	rows := mysql.conn.FetchRows(query, id)

	if rows == nil {
		return customer, fmt.Errorf("No se encontraron datos")
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&customer.Id, &customer.Name, &customer.Last_name, &customer.Phone_number, &customer.Curp, &customer.Number_license)
		if err != nil {
			return customer, err
		}
		return customer, nil
	}

	return customer, fmt.Errorf("No se encontró ningún cliente con el ID proporcionado")
}