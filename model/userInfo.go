package model

type UserInfo struct {
	Id           string   `json:"id" db:"id"`
	Name         string   `json:"name" db:"name"`
	Pickup       Location `json:"Pick_up"`
	Drop         Location `json:"Drop"`
	CabId        int32    `json:"CabId"`
	TravelTime   int64    `json:"TravelTime"`
	IsDisabled   bool     `json:"is_disabled" db:"is_disabled"`
	MobileNumber string   `json:"mobile_number" db:"mobile_number"`
	Token        string   `json:"token" db:"token"`
}

type UserPickup struct {
		Pickup     Location `json:"pick_up"`
		Drop       Location `json:"drop"`
}
