package dbconnect

import(
	"fmt"

	"database/sql"
	_ "github.com/lib/pq"
)

const (
	host = "localhost"
	port = 5432
	user = "postgres"
	password = "suseno10"
	dbname = "enigma_laundry"
)

var infoPsql = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

func ConnectDB() *sql.DB{
	db, err := sql.Open("postgres",infoPsql)
	if err != nil{
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}