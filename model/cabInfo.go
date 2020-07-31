package model

type CabInfo struct {
	ID                       string      `json:"id" db:"id"`
	Location                 Location    `json:"Location"`
	LastCabLocationLatitude  float64     `json:"last_cab_location_latitude" db:"last_cab_location_latitude"`
	LastCabLocationLongitude float64     `json:"last_cab_location_longitude" db:"last_cab_location_longitude"`
	IsAvailable              bool        `json:"IsAvailable" db:"is_available"`
	VehicleNumber            string      `json:"vehicle_number" db:"vehicle_number"`
	Distance                 float64
}
