package models

type UnitData struct {
	ID          int64  `json:"id"`
	UnitCode    string `json:"unitCode"`
	UnitType    string `json:"unitType"`
	Description string `json:"description"`
	Name        string `json:"name"`
}
