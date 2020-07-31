package model

type ResponseCabInfo struct {
	Location    Location           `json:"Location"`
	ID          string              `json:"id" db:"id"`
	Distance              float64
}
