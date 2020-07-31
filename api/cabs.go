package api

import(
  "time"
  "fmt"
  "github.com/gocql/gocql"
  "math/rand"
  models "cabservice/model"
  db "cabservice/db"
  "github.com/gin-gonic/gin"
  utils"cabservice/utils"
)


func InitializeCabs() {
	// Setting up my cabs here. Assuming a dozen cabs. This can be taken as an input too. For the sake
	// of convenience am setting it up here
  rand.Seed(time.Now().UnixNano())
  for rangenumber:= 0; rangenumber < 12; rangenumber ++ {
    latitde,longitude := getRandomLocationCoordinates(18.4529, 18.6298, 73.7389, 73.9787)

    cabInfo := models.CabInfo{
      ID : gocql.UUID.String(gocql.TimeUUID()),
      IsAvailable: true,
      LastCabLocationLatitude : latitde,
      LastCabLocationLongitude: longitude}
      addCabError := db.AddCabDetails(cabInfo)
      if addCabError != nil {
        fmt.Printf("\n\n Failed to Initialize Cab %+v due to %+v\n", cabInfo, addCabError)
      }
  }
}

func getRandomLocationCoordinates(latmin, latmax float64, longmin, longmax float64) (float64,float64) {

	randomLatitudeCoordinates := latmin + rand.Float64()*(latmax-latmin)
	randomLongitudeCoordinates := longmin + rand.Float64()*(longmax-longmin)

	return randomLatitudeCoordinates,randomLongitudeCoordinates
}



func CabV1Routes(router *gin.Engine) {

	cabRouter := router.Group("/v1/cab")
	{
    cabRouter.Use(ValidateUserRequestAndFetchUser())

		cabRouter.PUT("/endride/:id", endRide)
  }

}

func endRide(c *gin.Context){
  cabId := c.Param("id")
  fmt.Printf("\n cab id %+v", cabId)
  userId := c.GetString("user_id")
  fmt.Printf("\n user id %+v", userId)
  toggleCabAvailabilityStatusErr := db.ToggleCabAvailabilityStatus(cabId, true)
  if toggleCabAvailabilityStatusErr != nil {
    utils.PrintStackTrace("\n endRide 1 %+v", toggleCabAvailabilityStatusErr)
    utils.SendError(c, 200, utils.APP_ERROR_SERVER)
    return
  }else{
    var timeDiff int
    userRideDetails, userRideDetailsErr := db.GetUserRideDetailsByUserAndCabId(userId, cabId)
    if userRideDetailsErr != nil {
      utils.PrintStackTrace("\n endRide 2 %+v", userRideDetailsErr)
      timeDiff = 0
    }else{
      timeDiff = int(time.Now().Sub(userRideDetails.StartTime).Hours())
    }

    UpdateEndRideDetailsForUserErr := db.UpdateEndRideDetailsForUser(userId, cabId, time.Now(), timeDiff, "Finished")
    if UpdateEndRideDetailsForUserErr != nil {
      utils.PrintStackTrace("\n endRide 3 %+v", UpdateEndRideDetailsForUserErr)
      utils.SendError(c, 200, utils.APP_ERROR_SERVER)
      return
    }else{
      utils.SendSuccess(c, nil)
    }
  }
}
