package core

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
    _ "github.com/go-sql-driver/mysql"
)


type Conn_MySQL struct {
    DB *sql.DB
    Err string
}

func GetDBPool() *Conn_MySQL {
    error := ""
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error al cargar el archivo .env: %v", err)
    }

    dbHost:= os.Getenv("DB_HOST")
    dbUser := os.Getenv("DB_USERNAME")
    dbPass := os.Getenv("DB_PASSWORD")
    dbSchema := os.Getenv("DB_DATABASE")

    dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", dbUser, dbPass, dbHost, dbSchema)
    

    db, err := sql.Open("mysql", dsn)

    if err != nil {
        return &Conn_MySQL{Err: fmt.Sprintf("Error al abrir la base de datos: %v", err)}
    }

    // Configuramos el pool de conexiones
    db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

    // Para probar la conexión
    if err := db.Ping(); err != nil {
        db.Close()
        return &Conn_MySQL{Err: fmt.Sprintf("Error al verificar la conexión de la base de datos: %v", err)}
    }

    fmt.Println("Conexión a la base de datos exitosamente")

    return &Conn_MySQL{DB: db, Err: error}

}


func (conn *Conn_MySQL) ExecutePreparedQuery(query string, values ...interface{}) (sql.Result, error) {
	stmt, err := conn.DB.Prepare(query)
	if err != nil {
		return nil, fmt.Errorf("error al preparar la consulta: %w", err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(values...)
	if err != nil {
		return nil, fmt.Errorf("error al ejecutar la consulta preparada: %w", err)
	}

	return result, nil
}


func (conn *Conn_MySQL) FetchRows(query string, values ...interface{}) (*sql.Rows) {
	rows, err := conn.DB.Query(query, values...)
	if err != nil {
		fmt.Printf ("error al ejecutar la consulta SELECT: %w", err)
	}

	return rows
}

