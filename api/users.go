package api

import(
  "fmt"
  "sync"
  "math"
  "github.com/gin-gonic/gin"
   models"cabservice/model"
    db"cabservice/db"
   utils"cabservice/utils"
)

func UserCabV1Routes(router *gin.Engine) {

	userRouter := router.Group("/v1/user")
	{
		userRouter.GET("/book", bookCab)
  }

}

func bookCab(c *gin.Context){
  userInfo := models.UserInfo{}

  bindErr := c.BindJSON(&userInfo)
  if bindErr != nil {
    utils.SendError(c, 200, utils.APP_ERROR_INVALID_PAYLOAD)
		return
  }

  currentlyAvailableCabs, currentlyAvailableCabsErr := db.GetAvailableCabs(true)
  if currentlyAvailableCabsErr != nil{
    fmt.Printf("bookCab 1", currentlyAvailableCabsErr)
    utils.SendError(c, 200, utils.APP_ERROR_SERVER)
    return
  }else{
    lengthOfCabs := len(currentlyAvailableCabs)
    if lengthOfCabs == 0{
      utils.SendError(c, 200, utils.APP_ERROR_NO_CABS_AVAILABLE)
      return
    }else{
      var wg sync.WaitGroup
      for index, _ := range currentlyAvailableCabs {
        wg.Add(1)
        go CalculateDistance(&wg, &currentlyAvailableCabs, index, userInfo.Pickup.Latitude, userInfo.Pickup.Longitude)
      }
      wg.Wait()
		  quicksort(currentlyAvailableCabs, 0, lengthOfCabs-1)
      utils.SendSuccess(c, currentlyAvailableCabs[0])
    }
  }
}

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

func swap(a *models.CabInfo, b *models.CabInfo) {
	temp := *a
	*a = *b
	*b = temp
}

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
}

func hsin(theta float64) float64 {
	return math.Pow(math.Sin(theta/2), 2)
}
