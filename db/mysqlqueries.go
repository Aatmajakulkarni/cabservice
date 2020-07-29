package db

import (
	"fmt"
//	"time"

	// "github.com/go-sql-driver/mysql"
)

const (
	//mysql queries as constants

	//branch related queries
	ADD_CAB_DETAILS_QUERY string = `INSERT INTO cabs.cabs_info(id, last_cab_location_latitude, last_cab_location_longitude, is_available)VALUES (:id, :last_cab_location_latitude, :last_cab_location_longitude,
														 :is_available)`

	GET_ALL_AVAILABLE_CABS string = `SELECT cabs_info.id ,cabs_info.last_cab_location_latitude ,cabs_info.last_cab_location_longitude,cabs_info.is_available FROM cabs.cabs_info WHERE is_available=? `

	// mysql  tables declaration
	CREATE_DATABASE string = `CREATE DATABASE IF NOT EXISTS cabs DEFAULT CHARSET=utf8mb4;`

	TABLE_USER_RIDES string = "cabs.user_rides"
	TABLE_CABS_INFO string = "cabs.cabs_info"


	CREATE_TABLE_CABS_INFO string = `CREATE TABLE IF NOT EXISTS cabs_info(
                                 id varchar(255) PRIMARY KEY,
 																 last_cab_location_latitude float,
                                 last_cab_location_longitude float,
                                 is_available bool)
															 	 DEFAULT CHARSET=utf8mb4`


	CREATE_TABLE_USER_RIDES string = `CREATE TABLE IF NOT EXISTS user_rides(
                                id varchar(255) PRIMARY KEY,
                                name varchar(255),
                                pick_up_location_latitude float,
                                pick_up_location_longitude float,
                                drop_location_latitude float,
                                drop_location_longitude float,
                                cab_id varchar(255),
                                travel_time int,
                                mobile_number varchar(255),
														  	FOREIGN KEY(cab_id) REFERENCES cabs.cabs_info(id))
															  DEFAULT CHARSET=utf8mb4`


)

func init() {

	if _, err := mysqlDb.Exec(CREATE_DATABASE); nil != err {
		fmt.Printf("cannot create  CREATE_DATABASE%+v", err)
	}

	if _, err := mysqlDb.Exec(CREATE_TABLE_CABS_INFO); nil != err {
		fmt.Printf("cannot create  CREATE_TABLE_CABS_INFO %+v\n", err)
	}

	if _, err := mysqlDb.Exec(CREATE_TABLE_USER_RIDES); nil != err {
		fmt.Printf("cannot create  CREATE_TABLE_USER_RIDES %+v", err)
	}

}
