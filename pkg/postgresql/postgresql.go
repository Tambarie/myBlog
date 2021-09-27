package postgresql

import (
	"database/sql"
	"fmt"
)

func Db() *sql.DB {
	const (
		host = "localhost"
		port =5432
		user = "decagon"
		dbname = "tambarie"
	)

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable",host,port,user,dbname)

}
