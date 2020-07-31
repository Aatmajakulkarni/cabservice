package utils
const (
	SEVERITY_ERROR               = 0xC000 //49152
	FACILITIES_APPLICATION       = 0x0B00 //256
	FACILITIES_USER              = 0x0200 //512
	SUCCESS                      = 0

	APP_ERROR_SERVER                = SEVERITY_ERROR + FACILITIES_APPLICATION + 1  //49409
	APP_ERROR_INVALID_PAYLOAD       = SEVERITY_ERROR + FACILITIES_APPLICATION + 2  //49410
	APP_ERROR_FORCE_EXIT            = SEVERITY_ERROR + FACILITIES_APPLICATION + 3  //49411
	APP_ERROR_SOMETHING_WENT_WRONG  = SEVERITY_ERROR + FACILITIES_APPLICATION + 4  //49412
	APP_ERROR_FAILED_TO_BOOK_CAB    = SEVERITY_ERROR + FACILITIES_APPLICATION + 5  //49413
	APP_ERROR_FAILED_TO_END_RIDE    = SEVERITY_ERROR + FACILITIES_APPLICATION + 6  //49414
	APP_ERROR_INVALID_BOOKING       = SEVERITY_ERROR + FACILITIES_APPLICATION + 7  //49415
	APP_ERROR_ALREADY_EXISTS        = SEVERITY_ERROR + FACILITIES_APPLICATION + 8  //49416
	APP_ERROR_TIME_ERROR            = SEVERITY_ERROR + FACILITIES_APPLICATION + 9  //49417
	APP_ERROR_NO_CABS_AVAILABLE     = SEVERITY_ERROR + FACILITIES_APPLICATION + 10  //49418
	USER_ERROR_INVALID_SESSION      = SEVERITY_ERROR + FACILITIES_APPLICATION + 11  //49419
	USER_ERROR_NO_SUCH_USER         = SEVERITY_ERROR + FACILITIES_APPLICATION + 12  //49420

)

func GetErrorMessage(errorCode int) string {
	var message string
	switch errorCode {
	case APP_ERROR_SERVER:
		message = "Server Error"
	case APP_ERROR_INVALID_PAYLOAD:
		message = "Invalid parameters"
	case APP_ERROR_FORCE_EXIT:
		message = "Server forced exit"
  case APP_ERROR_SOMETHING_WENT_WRONG:
    message = "Sorry something went wrong"
  case APP_ERROR_FAILED_TO_BOOK_CAB:
    message = "Failed to book a cab"
  case APP_ERROR_FAILED_TO_END_RIDE:
    message = "Failed to end ride"
  case APP_ERROR_INVALID_BOOKING:
    message = "Invalid booking"
	case APP_ERROR_ALREADY_EXISTS:
		message = "Already Exists "
	case APP_ERROR_TIME_ERROR:
		message = "invalid time"
	case APP_ERROR_NO_CABS_AVAILABLE:
		message = "Sorry! no cabs are available currently"
	case USER_ERROR_INVALID_SESSION:
		message = "Invalid session"
	case USER_ERROR_NO_SUCH_USER:
		message = "No such user"
	}

	return message
}

func GetError(errorCode int) Response {
	message := GetErrorMessage(errorCode)
	return Response{Status: Status{StatusCode: errorCode, Message: message}}
}
