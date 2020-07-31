package api

import(
  "fmt"
  "sync"
  "math"
  "time"
  "github.com/gocql/gocql"
  "github.com/gin-gonic/gin"
  models"cabservice/model"
  db"cabservice/db"
  utils"cabservice/utils"
)

/*
Function name : UserCabV1Routes
Description  : router group for user REST api
Params       :  *gin.Engine
Return       :  nil
*/

func UserCabV1Routes(router *gin.Engine) {

  loginRouter := router.Group("v1/new")
  {
    //for new user login
    loginRouter.POST("login", newUserLogin)
  }

	userRouter := router.Group("/v1/user")
	{
    userRouter.Use(ValidateUserRequestAndFetchUser())

    //user can book a cab
    userRouter.POST("/book", bookCab)

    //user's all past rides
    userRouter.GET("/rides", getUserRides)
  }

}

/*
Function name : newUserLogin
Description  :  this function is for new user login if the user already exists then it returns token for existing user
                or creates new encrypted token with encryption key
Params       :  c *gin.Context
Return       :  nil
*/

func newUserLogin(c *gin.Context){
	var userInfo models.UserInfo
	c.BindJSON(&userInfo)


  fmt.Printf("\n user existring %+v %+v", userInfo.Name, userInfo.MobileNumber)
  existingUser, existingUserErr := db.GetUserInfo(userInfo.Name, userInfo.MobileNumber)
  if existingUserErr == nil {
    fmt.Printf("\n user existring")
    payload := models.PayloadLogin{
  		Token:        existingUser.Token,
  		UserInfo:     existingUser}
  	utils.SendSuccess(c, payload)
    return
  }
  fmt.Printf("\n error %+v",existingUserErr )
  userInfo.Id = gocql.UUID.String(gocql.TimeUUID())
  userInfo.IsDisabled = false
  userInfo.Token = utils.CreateToken(userInfo.Id)
  createUserErr := db.CreateUser(userInfo)
  if createUserErr != nil {
    fmt.Printf("\n newUserLogin %+v", createUserErr)
    utils.SendError(c, 200, utils.APP_ERROR_SERVER)
    return
  }
	// Create session token for the user token
	payload := models.PayloadLogin{
		Token:        userInfo.Token,
		UserInfo:     userInfo}
	utils.SendSuccess(c, payload)
}

/*
Function name : bookCab
Description  :  this function fetches all available cabs and calculates Distance for each cab from user's booking location and sends
                nearest cab details as a reponse
Params       :  c *gin.Context
Return       :  nil
*/

func bookCab(c *gin.Context){
  var pickup models.UserPickup

  bindErr := c.BindJSON(&pickup)
  if bindErr != nil {
    fmt.Printf("\n findCab 1 %+v\n", bindErr)
    utils.SendError(c, 200, utils.APP_ERROR_INVALID_PAYLOAD)
		return
  }

  user, ok := getUserFromSession(c)
  if ok {
  //fmt.Printf("\n data binded successfully\n")
  currentlyAvailableCabs, currentlyAvailableCabsErr := db.GetAvailableCabs(true)
  if currentlyAvailableCabsErr != nil{
    fmt.Printf("bookCab 1", currentlyAvailableCabsErr)
    utils.SendError(c, 200, utils.APP_ERROR_SERVER)
    return
  }else{
    fmt.Printf("\n available cabs %+v\n", len(currentlyAvailableCabs))
    lengthOfCabs := len(currentlyAvailableCabs)
    if lengthOfCabs == 0{
      utils.SendError(c, 200, utils.APP_ERROR_NO_CABS_AVAILABLE)
      return
    }else{
      var wg sync.WaitGroup
      for index, _ := range currentlyAvailableCabs {
        wg.Add(1)
        go CalculateDistance(&wg, &currentlyAvailableCabs, index, user.Pickup.Latitude, user.Pickup.Longitude)
      }
      wg.Wait()
		  quicksort(currentlyAvailableCabs, 0, lengthOfCabs-1)
      responseCabInfo := models.ResponseCabInfo{
        ID : currentlyAvailableCabs[0].ID,
        Location : models.Location{
          Latitude :currentlyAvailableCabs[0].LastCabLocationLatitude,
          Longitude:currentlyAvailableCabs[0].LastCabLocationLongitude},
          Distance: currentlyAvailableCabs[0].Distance}
    toggleCabAvailabilityStatusErr := db.ToggleCabAvailabilityStatus(currentlyAvailableCabs[0].ID, false)
      if toggleCabAvailabilityStatusErr != nil {
        utils.PrintStackTrace("\n bookCab 2 %+v", toggleCabAvailabilityStatusErr)
        utils.SendError(c, 200, utils.APP_ERROR_SERVER)
        return
      }else{
        rideDetails := models.UserRideInfo{
          ID : gocql.UUID.String(gocql.TimeUUID()),
          UserId: user.Id,
          CabId: currentlyAvailableCabs[0].ID,
          PickupLocationLatitude: pickup.Pickup.Latitude,
          PickupLocationLongitude: pickup.Pickup.Longitude,
          DropLocationLatitude: pickup.Drop.Latitude,
          DropLocationLongitude: pickup.Drop.Longitude,
          StartTime: time.Now(),
          TravelTime:0,
          RideStatus: "Booked"}
        addUserRideDetailsErr := db.AddUserRideDetails(rideDetails)
        if addUserRideDetailsErr != nil {
          utils.PrintStackTrace("\n bookCab 3 %+v", addUserRideDetailsErr)
          utils.SendError(c, 200, utils.APP_ERROR_SERVER)
          return
        }else{
          //fmt.Printf("\n nearest cab details%+v\n", responseCabInfo)
          utils.SendSuccess(c, responseCabInfo)
        }
      }

      }
    }
  }else {
		utils.SendError(c, 200, utils.USER_ERROR_NO_SUCH_USER)
	}
}

/*
Function name : quicksort
Description  :  this function uses quick sort from cabs Distances
Params       :  []models.CabInfo, leftIndex, rightIndex
Return       :  nil
*/

func quicksort(result []models.CabInfo, leftIndex int, rightIndex int) {

	if leftIndex >= rightIndex {
		return
	}
	pivot := result[rightIndex].Distance

	cnt := leftIndex

	for i := leftIndex; i <= rightIndex; i++ {

		if result[i].Distance <= pivot {
			swap(&result[cnt], &result[i])
			cnt++
		}
	}
	quicksort(result, leftIndex, cnt-2)
	quicksort(result, cnt, rightIndex)
}

/*
Function name : swap
Description  :  this function interchanges cabInfo
Params       :  *models.CabInfo, *models.CabInfo
Return       :  nil
*/

func swap(a *models.CabInfo, b *models.CabInfo) {
	temp := *a
	*a = *b
	*b = temp
}

/*
Function name : CalculateDistance
Description  :  this function calculate Distance between 2 locations
Params       :  pointer to WaitGroup, array of available cabs, current index, userlatitude, userlongitude
Return       :  nil
*/


func CalculateDistance(wg *sync.WaitGroup, currentlyAvailableCabs *[]models.CabInfo, index int, userLat float64, userLong float64) {

	// convert to radians
	// must cast radius as float to multiply later
	var la1, lo1, la2, lo2, r float64
	la1 = userLat * math.Pi / 180
	lo1 = userLong * math.Pi / 180
	la2 = (*currentlyAvailableCabs)[index].LastCabLocationLatitude * math.Pi / 180
	lo2 = (*currentlyAvailableCabs)[index].LastCabLocationLongitude * math.Pi / 180

	r = 6378100 // Earth radius in METERS

	// calculate
	h := hsin(la2-la1) + math.Cos(la1)*math.Cos(la2)*hsin(lo2-lo1)

	(*currentlyAvailableCabs)[index].Distance = 2 * r * math.Asin(math.Sqrt(h))
	//fmt.Printf("\ndistance 1 %+v\n", a)
  wg.Done()
}

func hsin(theta float64) float64 {
	return math.Pow(math.Sin(theta/2), 2)
}

/*
Function name : getUserRides
Description  :  this function returns all past user rides
Params       :  c *gin.Context
Return       :  nil
*/

func getUserRides(c *gin.Context){
  userId := c.GetString("user_id")

  rideDetailsForUser, rideDetailsForUserErr := db.GetRideDetailsForUser(userId)
  if rideDetailsForUserErr != nil {
    utils.PrintStackTrace("\n getUserRides 1 %+v", rideDetailsForUserErr)
    utils.SendError(c, 200, utils.APP_ERROR_SERVER)
    return
  }else{
    //fmt.Printf("\n nearest cab details%+v\n", responseCabInfo)
    utils.SendSuccess(c, rideDetailsForUser)
  }
}
