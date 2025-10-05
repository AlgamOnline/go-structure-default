package controllers

import (
	"encoding/json"
	"golang-default/models"
	"golang-default/services"
	"golang-default/utils"
	"net/http"
)

func GPSHandler(gpsService *services.GPSService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var data models.GPSData
		if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
			http.Error(w, "Invalid GPS data", http.StatusBadRequest)
			return
		}

		_, err := gpsService.InsertGPS(data)
		if err != nil {
			utils.JSON(w, http.StatusInternalServerError, false, nil, err.Error())
			return
		}

		err = gpsService.InsertLastGPS(data)
		if err != nil {
			utils.JSON(w, http.StatusInternalServerError, false, nil, err.Error())
			return
		}

		utils.JSON(w, http.StatusCreated, true, "", "GPS created successfully")
	}
}
