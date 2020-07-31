package db
import(
	models "cabservice/model"
	utils"cabservice/utils"
)

/*
Function name : AddCabDetails
Description  :  this function adds cab details in mysql table
Params       :  models.CabInfo
Return       :  error
*/

func AddCabDetails(cabDetails models.CabInfo) error {

	_, err := mysqlDb.NamedExec(ADD_CAB_DETAILS_QUERY, cabDetails)
	if err != nil {
		utils.PrintStackTrace("AddCabDetails 1", err)
	}
	return err
}

/*
Function name : GetAvailableCabs
Description  :  this function gives available cabs from mysql table
Params       :  IsAvailable
Return       :  []models.CabInfo, error
*/

func GetAvailableCabs(IsAvailable bool)([]models.CabInfo, error){

	cabsRecords := []models.CabInfo{}
	cabsRecordsErr := mysqlDb.Select(&cabsRecords, GET_ALL_AVAILABLE_CABS, IsAvailable)

	if cabsRecordsErr != nil {
		utils.PrintStackTrace("GetAvailableCabs ", cabsRecordsErr)
	}
	return cabsRecords, cabsRecordsErr
}

/*
Function name : ToggleCabAvailabilityStatus
Description  :  this function changes availability status for given cab
Params       :  cabId, status
Return       :  error
*/

func ToggleCabAvailabilityStatus(cabId string, status bool) error{
	_, err := mysqlDb.Exec(UPDATE_CAB_AVAILABILITY_STATUS_QUERY, status, cabId)
	if err != nil {
		utils.PrintStackTrace("ToggleCabAvailabilityStatus 1", err)
	}
	return err
}
