package model

type UserInfo struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Pickup     Location `json:"Pick_up"`
	Drop       Location `json:"Drop"`
	CabId      int32             `json:"CabId"`
	TravelTime int64             `json:"TravelTime"`
}
