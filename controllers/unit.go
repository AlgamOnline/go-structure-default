package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"golang-default/models"
	"golang-default/services"
	"golang-default/utils"

	"github.com/gorilla/mux"
)

func CreateUnitHandler(UnitService *services.UnitService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var unit models.UnitData
		if err := json.NewDecoder(r.Body).Decode(&unit); err != nil {
			utils.JSON(w, http.StatusBadRequest, false, nil, "Invalid request body")
			return
		}

		id, err := UnitService.CreateUnit(unit)
		if err != nil {
			utils.JSON(w, http.StatusInternalServerError, false, nil, err.Error())
			return
		}

		unit.ID = id
		utils.JSON(w, http.StatusCreated, true, unit, "Unit created successfully")
	}
}

func GetUnitHandler(UnitService *services.UnitService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := mux.Vars(r)["id"]
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			utils.JSON(w, http.StatusBadRequest, false, nil, "Invalid Unit ID")
			return
		}

		user, err := UnitService.GetUnitByID(id)
		if err != nil {
			utils.JSON(w, http.StatusNotFound, false, nil, "Unit not found")
			return
		}

		utils.JSON(w, http.StatusOK, true, user, "Unit fetched successfully")
	}
}

func UpdateUnitHandler(UnitService *services.UnitService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := mux.Vars(r)["id"]
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			utils.JSON(w, http.StatusBadRequest, false, nil, "Invalid user ID")
			return
		}

		var user models.UnitData
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			utils.JSON(w, http.StatusBadRequest, false, nil, "Invalid request body")
			return
		}
		user.ID = id

		if err := UnitService.UpdateUnit(user); err != nil {
			utils.JSON(w, http.StatusInternalServerError, false, nil, err.Error())
			return
		}

		utils.JSON(w, http.StatusOK, true, user, "Unit updated successfully")
	}
}

func DeleteUnitHandler(UnitService *services.UnitService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := mux.Vars(r)["id"]
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			utils.JSON(w, http.StatusBadRequest, false, nil, "Invalid user ID")
			return
		}

		if err := UnitService.DeleteUnit(id); err != nil {
			utils.JSON(w, http.StatusInternalServerError, false, nil, err.Error())
			return
		}

		utils.JSON(w, http.StatusOK, true, nil, "Unit deleted successfully")
	}
}
