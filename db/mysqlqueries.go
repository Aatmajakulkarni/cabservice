package db

import (
	"fmt"
//	"time"

	// "github.com/go-sql-driver/mysql"
)

const (
	//mysql queries as constants

	//branch related queries
	ADD_CAB_DETAILS_QUERY string = `INSERT INTO ` + TABLE_CABS_INFO + `(id, last_cab_location_latitude, last_cab_location_longitude, is_available)VALUES (:id, :last_cab_location_latitude, :last_cab_location_longitude,
														 :is_available)`

	GET_ALL_AVAILABLE_CABS string = `SELECT cabs_info.id ,cabs_info.last_cab_location_latitude ,cabs_info.last_cab_location_longitude,cabs_info.is_available FROM ` + TABLE_CABS_INFO + ` WHERE is_available=? `

	GET_USER_INFO_BY_ID_QUERY string = `SELECT user_info.id ,user_info.name ,user_info.is_disabled, user_info.mobile_number FROM ` + TABLE_USER_INFO + ` WHERE id=? `

	GET_USER_INFO_QUERY string = `SELECT user_info.id ,user_info.name ,user_info.is_disabled,user_info.mobile_number, user_info.token FROM ` + TABLE_USER_INFO + ` WHERE name=? and mobile_number=?`

	ADD_USER_INFO_QUERY string = `INSERT INTO ` + TABLE_USER_INFO + `(id, name, is_disabled, mobile_number, token)VALUES (:id, :name, :is_disabled, :mobile_number, :token)`

	UPDATE_CAB_AVAILABILITY_STATUS_QUERY string = `update ` +  TABLE_CABS_INFO + ` set is_available=? where id=?`

	ADD_USER_RIDE_DETAILS_QUERY string = `INSERT INTO ` + TABLE_USER_RIDES + `(id, user_id, pick_up_location_latitude, pick_up_location_longitude, drop_location_latitude, drop_location_longitude, cab_id, travel_time, start_time, ride_status)VALUES (:id, :user_id, :pick_up_location_latitude, :pick_up_location_longitude, :drop_location_latitude, :drop_location_longitude, :cab_id, :travel_time, :start_time, :ride_status)`

  GET_USER_RIDE_DETAILS_QUERY string = `SELECT user_rides.id, user_rides.user_id, user_rides.pick_up_location_latitude, user_rides.pick_up_location_longitude, user_rides.drop_location_latitude, user_rides.drop_location_longitude, user_rides.cab_id, user_rides.travel_time, user_rides.start_time, user_rides.ride_status from ` + TABLE_USER_RIDES + ` where user_id=? and cab_id=?`

	UPDATE_RIDE_END_DATA_FOR_USER_RIDE_QUERY string = `update ` +  TABLE_USER_RIDES + ` set start_time=?, travel_time=?, ride_status=? where user_id=? and cab_id=?`

	GET_USER_RIDES_DETAILS_QUERY string = `SELECT user_rides.id, user_rides.user_id, user_rides.pick_up_location_latitude, user_rides.pick_up_location_longitude, user_rides.drop_location_latitude, user_rides.drop_location_longitude, user_rides.cab_id, user_rides.travel_time, user_rides.start_time, user_rides.ride_status from ` + TABLE_USER_RIDES + ` where user_id=?`

	// mysql  tables declaration
	CREATE_DATABASE string = `CREATE DATABASE IF NOT EXISTS cabs DEFAULT CHARSET=utf8mb4;`

	TABLE_USER_RIDES string = "cabs.user_rides"
	TABLE_CABS_INFO string = "cabs.cabs_info"
	TABLE_USER_INFO string = "cabs.user_info"


	CREATE_TABLE_CABS_INFO string = `CREATE TABLE IF NOT EXISTS cabs_info(
                                 id varchar(255) PRIMARY KEY,
 																 last_cab_location_latitude double,
                                 last_cab_location_longitude double,
																 vehicle_number varchar(255),
                                 is_available bool)
															 	 DEFAULT CHARSET=utf8mb4`


	CREATE_TABLE_USER_RIDES string = `CREATE TABLE IF NOT EXISTS user_rides(
                                id varchar(255) PRIMARY KEY,
																user_id varchar(255),
                                pick_up_location_latitude double,
                                pick_up_location_longitude double,
                                drop_location_latitude double,
                                drop_location_longitude double,
                                cab_id varchar(255),
                                travel_time int,
																start_time timestamp,
																end_timestamp time DEFAULT '1970-01-01 00:00:00',
																ride_status varchar(255),
														  	FOREIGN KEY(cab_id) REFERENCES cabs.cabs_info(id),
																FOREIGN KEY(user_id) REFERENCES cabs.user_info(id))
															  DEFAULT CHARSET=utf8mb4`

	CREATE_TABLE_USER_INFO string = `CREATE TABLE IF NOT EXISTS user_info(
                                 id varchar(255) PRIMARY KEY,
 																 name varchar(255),
																 is_disabled bool,
																 token text,
                                 mobile_number varchar(255))
															 	 DEFAULT CHARSET=utf8mb4`


)

func init() {

	if _, err := mysqlDb.Exec(CREATE_DATABASE); nil != err {
		fmt.Printf("cannot create  CREATE_DATABASE%+v", err)
	}

	if _, err := mysqlDb.Exec(CREATE_TABLE_CABS_INFO); nil != err {
		fmt.Printf("cannot create  CREATE_TABLE_CABS_INFO %+v\n", err)
	}

	if _, err := mysqlDb.Exec(CREATE_TABLE_USER_INFO); nil != err {
		fmt.Printf("cannot create  CREATE_TABLE_USER_INFO %+v", err)
	}

	if _, err := mysqlDb.Exec(CREATE_TABLE_USER_RIDES); nil != err {
		fmt.Printf("cannot create  CREATE_TABLE_USER_RIDES %+v", err)
	}

}
