package model

type CabInfo struct {
	Location    Location           `json:"Location"`
	LastCabLocationLatitude float64  `json:"last_cab_location_latitude" db:"last_cab_location_latitude"`
	LastCabLocationLongitude float64  `json:"last_cab_location_longitude" db:"last_cab_location_longitude"`
	IsAvailable bool               `json:"IsAvailable" db:"is_available"`
	ID          string              `json:"id" db:"id"`
	Distance              float64
}
