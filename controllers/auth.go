package controllers

import (
	"encoding/json"
	"golang-default/models"
	"golang-default/services"
	"golang-default/utils"
	"net/http"
)

func LoginHandler(authService *services.AuthService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.LoginRequest
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			utils.JSON(w, http.StatusBadRequest, false, nil, "Invalid request body")
			return
		}

		data, err := authService.Login(user.Email, user.Password)
		if err != nil {
			utils.JSON(w, http.StatusInternalServerError, false, nil, err.Error())
			return
		}

		utils.JSON(w, http.StatusCreated, true, data, "User Login successfully")
	}
}
