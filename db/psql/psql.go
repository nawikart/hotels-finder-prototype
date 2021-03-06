package psql

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// const (
// 	host     = "db-postgresql-sgp1-agoda-do-user-4803185-0.db.ondigitalocean.com"
// 	port     = 25060
// 	user     = "doadmin"
// 	password = "btz4gefi370myn3z"
// 	dbname   = "defaultdb"
// )

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "root"
	dbname   = "hotels"
)

func Connect() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return db, nil
}
