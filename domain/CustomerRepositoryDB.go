package domain

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type CustomerRepositoryDb struct {
	client *sql.DB
}

func (d CustomerRepositoryDb) FindAll() ([]Customer, error) {

	findAllSql := "select customer_id, name, city,zipcode, date_of_birth, status from customers"
	rows, err := d.client.Query(findAllSql)
	if err != nil {
		log.Println("error while qeurying customer table" + err.Error())
		return nil, err
	}

	customers := make([]Customer, 0)
	for rows.Next() {
		var c Customer
		err = rows.Scan(&c.Id, &c.Name, &c.City, &c.ZipCode, &c.DateOfBirth, &c.Status)
		if err != nil {
			log.Println("error while scanning customer table" + err.Error())
			return nil, err
		}
		customers = append(customers, c)
	}
	return customers, nil
}

func NewCustomerRepositoryDb() CustomerRepository {
	client, err := sql.Open("mysql", "root:@tcp(localhost:3306)/banking")
	if err != nil {
		panic(err)
	}

	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	return CustomerRepositoryDb{client}
}