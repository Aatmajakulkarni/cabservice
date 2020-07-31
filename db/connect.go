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

/*
Function name : init
Description  :  this init function initiate mysql connection for project
Params       :  takes MYSQL_SERVER Param value from os.(eg. in .bashrc file - export MYSQL_SERVER=root:root@tcp\(127.0.0.1\)/cabs?parseTime=true)
Return       :  nil
*/

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
