package db

import (
	"fmt"
	//"time"
  "os"
	_ "github.com/go-sql-driver/mysql"
	 "github.com/jmoiron/sqlx"
)

var (
	mysqlDb                 *sqlx.DB
)

func init() {

	//mysql server
	//	fmt.Println("envi", os.Getenv("CLIMB_MYSQL_SERVER"))
	mysql, err := sqlx.Open("mysql", (os.Getenv("MYSQL_SERVER")))

	if err != nil {
		panic(err.Error())
	} else {
		fmt.Printf("\n connection established")
		mysqlDb = mysql
	}
	//defer mysqlDb.Close()

}
