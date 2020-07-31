package db

import(
	"time"
	models "cabservice/model"
	utils"cabservice/utils"
)
func GetUserInfoById(userId string) (models.UserInfo, error) {
	userData := models.UserInfo{}
	err := mysqlDb.Get(&userData, GET_USER_INFO_BY_ID_QUERY, userId)

	return userData, err
}

func GetUserInfo(name string, mobileNumber string) (models.UserInfo, error) {
	userData := models.UserInfo{}
	err := mysqlDb.Get(&userData, GET_USER_INFO_QUERY, name, mobileNumber)

	return userData, err
}

func CreateUser(userInfo models.UserInfo) error {
	_, err := mysqlDb.NamedExec(ADD_USER_INFO_QUERY, userInfo)
	if err != nil {
		utils.PrintStackTrace("CreateUser 1", err)
	}
	return err

}

func AddUserRideDetails(rideDetails models.UserRideInfo) error {
	_, err := mysqlDb.NamedExec(ADD_USER_RIDE_DETAILS_QUERY, rideDetails)
	if err != nil {
		utils.PrintStackTrace("AddUserRideDetails 1", err)
	}
	return err
}

func GetUserRideDetailsByUserAndCabId(userId string, cabId string)(models.UserRideInfo, error){
	userRideDetails := models.UserRideInfo{}
	err := mysqlDb.Get(&userRideDetails, GET_USER_RIDE_DETAILS_QUERY, userId, cabId)

	return userRideDetails, err
}

func UpdateEndRideDetailsForUser(userId string, cabId string, startTime time.Time, travelTime int, status string) error {
	_, err := mysqlDb.Exec(UPDATE_RIDE_END_DATA_FOR_USER_RIDE_QUERY, startTime, travelTime, status, userId, cabId)
	if err != nil {
		utils.PrintStackTrace("UpdateEndRideDetailsForUser 1", err)
	}
	return err
}

func GetRideDetailsForUser(userId string) ([]models.UserRideInfo, error) {
	userRidesDetails := []models.UserRideInfo{}
	err := mysqlDb.Select(&userRidesDetails, GET_USER_RIDES_DETAILS_QUERY, userId)

	return userRidesDetails, err
}
