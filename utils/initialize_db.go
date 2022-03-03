package utils

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func Set_db(Db_conn_str string) *sql.DB {
	db, err := sql.Open("postgres", Db_conn_str)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Подключение к базе данных было успешно")
	}
	// defer db.Close()
	return db
}
