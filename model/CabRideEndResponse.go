package model

type CabRideEndResponse struct {
	Bill     int64  `json:"Bill"`
	Distance int64  `json:"Distance"`
	IsEnded  bool   `json:"IsEnded"`
	Errors   string `json:"Errors"`
}
