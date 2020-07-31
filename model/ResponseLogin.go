package model

type PayloadLogin struct {
	Token        string `json:"token"`
	UserInfo         UserInfo   `json:"user_info"`
}
