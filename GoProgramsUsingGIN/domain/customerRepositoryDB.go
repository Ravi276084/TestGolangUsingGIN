package domain

import (
	"database/sql"
	"log"
	"time"

	"github.com/Ravi276084/TestGoLangUsingGIN/GoProgramsUsingGIN/errs"
	_ "github.com/lib/pq"
)

type CustomerRepositoryDb struct {
	client *sql.DB
}

func (d CustomerRepositoryDb) FindAll() ([]Customer, *errs.AppError) {
	findAllSql := "select Id, Name, City, Zipcode, DateofBirth, Status from customers"

	rows, err := d.client.Query(findAllSql)
	if err != nil {
		log.Println("Error while querying customer table " + err.Error())
		return nil, errs.NewUnexpectedError("Database error")
	}

	customers := make([]Customer, 0)
	for rows.Next() {
		var c Customer
		err := rows.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateofBrith, &c.Status)
		if err != nil {
			log.Println("Error while scanning customers" + err.Error())
			return nil, errs.NewUnexpectedError("Database error")
		}
		customers = append(customers, c)
	}
	return customers, nil
}

func (d CustomerRepositoryDb) ById(id int) (*Customer, *errs.AppError) {
	customerSql := "select id, name, city, zipcode, dateofbirth, status from customers where id=$1"

	row := d.client.QueryRow(customerSql, id)
	var c Customer
	err := row.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateofBrith, &c.Status)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NotFoundError("Customer not found")
		} else {
			log.Println("Error while scanning customer " + err.Error())
			return nil, errs.NewUnexpectedError("Database error")
		}
	}
	return &c, nil
}

func NewCustomerRepositoryDb() CustomerRepositoryDb {
	dsn := "user=postgres password=Prasad@84 dbname=postgres sslmode=disable"
	client, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	// client, err := sql.Open("mysql", "postgres:Prasad@84@tcp(localhost:5432)/TestDataBase")
	// if err != nil {
	// 	panic(err)
	// }
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return CustomerRepositoryDb{client}
}
