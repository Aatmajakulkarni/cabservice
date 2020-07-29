package db
import(
	models "cabservice/model"
	utils"cabservice/utils"
)
func AddCabDetails(cabDetails models.CabInfo) error {

	_, err := mysqlDb.NamedExec(ADD_CAB_DETAILS_QUERY, cabDetails)
	if err != nil {
		utils.PrintStackTrace("AddCabDetails 1", err)
	}
	return err
}

func GetAvailableCabs(IsAvailable bool)([]models.CabInfo, error){

	cabsRecords := []models.CabInfo{}
	cabsRecordsErr := mysqlDb.Select(&cabsRecords, GET_ALL_AVAILABLE_CABS, IsAvailable)

	if cabsRecordsErr != nil {
		utils.PrintStackTrace("GetAvailableCabs ", cabsRecordsErr)
	}
	return cabsRecords, cabsRecordsErr
}
