package db

import(
	"time"
	models "cabservice/model"
	utils"cabservice/utils"
)

/*
Function name : GetUserInfoById
Description   : this functions returns user info by given user id
Params        : userId
Return        :  models.UserInfo, error
*/

func GetUserInfoById(userId string) (models.UserInfo, error) {
	userData := models.UserInfo{}
	err := mysqlDb.Get(&userData, GET_USER_INFO_BY_ID_QUERY, userId)

	return userData, err
}


/*
Function name : GetUserInfo
Description   : this functions returns user info by given user name, mobilenumber
Params        : name, mobileNumber
Return        :  models.UserInfo, error
*/

func GetUserInfo(name string, mobileNumber string) (models.UserInfo, error) {
	userData := models.UserInfo{}
	err := mysqlDb.Get(&userData, GET_USER_INFO_QUERY, name, mobileNumber)

	return userData, err
}

/*
Function name : CreateUser
Description   : this function enters user data in mysql table
Params        : models.UserInfo
Return        :  error
*/

func CreateUser(userInfo models.UserInfo) error {
	_, err := mysqlDb.NamedExec(ADD_USER_INFO_QUERY, userInfo)
	if err != nil {
		utils.PrintStackTrace("CreateUser 1", err)
	}
	return err

}

/*
Function name : AddUserRideDetails
Description   : this function adds ride details in table
Params        : models.UserRideInfo
Return        :  error
*/

func AddUserRideDetails(rideDetails models.UserRideInfo) error {
	_, err := mysqlDb.NamedExec(ADD_USER_RIDE_DETAILS_QUERY, rideDetails)
	if err != nil {
		utils.PrintStackTrace("AddUserRideDetails 1", err)
	}
	return err
}

/*
Function name : GetUserRideDetailsByUserAndCabId
Description   : this function returns ride details by  given userId, cabId
Params        : userid, cabId
Return        : models.UserRideInfo, error
*/

func GetUserRideDetailsByUserAndCabId(userId string, cabId string)(models.UserRideInfo, error){
	userRideDetails := models.UserRideInfo{}
	err := mysqlDb.Get(&userRideDetails, GET_USER_RIDE_DETAILS_QUERY, userId, cabId)

	return userRideDetails, err
}

/*
Function name : UpdateEndRideDetailsForUser
Description   : this function updates ride details in the database
Params        : userid, cabId, startTime, travelTime, status
Return        : error
*/

func UpdateEndRideDetailsForUser(userId string, cabId string, startTime time.Time, travelTime int, status string) error {
	_, err := mysqlDb.Exec(UPDATE_RIDE_END_DATA_FOR_USER_RIDE_QUERY, startTime, travelTime, status, userId, cabId)
	if err != nil {
		utils.PrintStackTrace("UpdateEndRideDetailsForUser 1", err)
	}
	return err
}

/*
Function name : GetRideDetailsForUser
Description   : this function returns all the user past rides by userId from database
Params        : userid
Return        : []models.UserRideInfo, error
*/

func GetRideDetailsForUser(userId string) ([]models.UserRideInfo, error) {
	userRidesDetails := []models.UserRideInfo{}
	err := mysqlDb.Select(&userRidesDetails, GET_USER_RIDES_DETAILS_QUERY, userId)

	return userRidesDetails, err
}
