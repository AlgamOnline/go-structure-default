package models

// Jangan import "golang-default/controllers" atau "golang-default/services" di sini
import "time" // contoh import yang boleh
type GPSData struct {
	Id        int64     `json:"id"`
	UnitId    int64     `json:"unitId"`
	Latitude  float64   `json:"latitude"`
	Longitude float64   `json:"longitude"`
	Speed     float64   `json:"speed"`
	Bearing   float64   `json:"bearing"`
	Accuracy  float64   `json:"accuracy"`
	Altitude  float64   `json:"altitude"`
	CreatedAt time.Time `json:"createdAt"`
}

type LastGPSData struct {
	UnitId    int64     `json:"unitId"`
	Latitude  float64   `json:"latitude"`
	Longitude float64   `json:"longitude"`
	Speed     float64   `json:"speed"`
	Bearing   float64   `json:"bearing"`
	Accuracy  float64   `json:"accuracy"`
	Altitude  float64   `json:"altitude"`
	UpdatedAt time.Time `json:"updatedAt"`
}
