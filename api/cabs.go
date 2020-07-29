package api

import(
  "time"
  "fmt"
  "github.com/gocql/gocql"
  "math/rand"
  models "cabservice/model"
  db "cabservice/db"
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
