package model

import(
	"time"
)

type UserRideInfo struct {
	ID                       string      `json:"id" db:"id"`
	UserId                   string      `json:"user_id" db:"user_id"`
	CabId                    string      `json:"cab_id" db:"cab_id"`
	Location                 Location    `json:"Location"`
	PickupLocationLatitude   float64     `json:"pick_up_location_latitude" db:"pick_up_location_latitude"`
	PickupLocationLongitude  float64     `json:"pick_up_location_longitude" db:"pick_up_location_longitude"`
	DropLocationLatitude     float64     `json:"drop_location_latitude" db:"drop_location_latitude"`
	DropLocationLongitude    float64     `json:"drop_location_longitude" db:"drop_location_longitude"`
	StartTime                time.Time   `json:"start_time" db:"start_time"`
	EndTime                  time.Time   `json:"end_time" db:"end_time"`
	TravelTime               int         `json:"travel_time" db:"travel_time"`
	RideStatus               string      `json:"ride_status" db:"ride_status"`
}
