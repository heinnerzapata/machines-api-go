package connect

import (
	"database/sql"
	"fmt"
	"machines-api-go/db"
	"os"

	_ "github.com/lib/pq"
)

var DataQueries *db.Queries
var DbQueries *sql.DB

func Connect() {
	var err error

	fmt.Println("Connecting DB...")
	DbQueries, err = sql.Open("postgres", os.Getenv("DB_STRING_CONN"))

	if err != nil {
		fmt.Println("Connecting DB error !")
		panic(err)
	}

	// defer DbQueries.Close()

	err = DbQueries.Ping()
	if err != nil {
		fmt.Println("panic connect !!")
		panic(err)
	}
	fmt.Println("Established a successful connection with DB!")

	DataQueries = db.New(DbQueries)
}
