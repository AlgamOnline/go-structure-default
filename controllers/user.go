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

func CreateUserHandler(userService *services.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.UserData
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			utils.JSON(w, http.StatusBadRequest, false, nil, "Invalid request body")
			return
		}

		id, err := userService.CreateUser(user)
		if err != nil {
			utils.JSON(w, http.StatusInternalServerError, false, nil, err.Error())
			return
		}

		user.ID = id
		utils.JSON(w, http.StatusCreated, true, user, "User created successfully")
	}
}

func GetUserHandler(userService *services.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := mux.Vars(r)["id"]
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			utils.JSON(w, http.StatusBadRequest, false, nil, "Invalid user ID")
			return
		}

		user, err := userService.GetUserByID(id)
		if err != nil {
			utils.JSON(w, http.StatusNotFound, false, nil, "User not found")
			return
		}

		utils.JSON(w, http.StatusOK, true, user, "User fetched successfully")
	}
}

func UpdateUserHandler(userService *services.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := mux.Vars(r)["id"]
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			utils.JSON(w, http.StatusBadRequest, false, nil, "Invalid user ID")
			return
		}

		var user models.UserData
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			utils.JSON(w, http.StatusBadRequest, false, nil, "Invalid request body")
			return
		}
		user.ID = id

		if err := userService.UpdateUser(user); err != nil {
			utils.JSON(w, http.StatusInternalServerError, false, nil, err.Error())
			return
		}

		utils.JSON(w, http.StatusOK, true, user, "User updated successfully")
	}
}

func DeleteUserHandler(userService *services.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := mux.Vars(r)["id"]
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			utils.JSON(w, http.StatusBadRequest, false, nil, "Invalid user ID")
			return
		}

		if err := userService.DeleteUser(id); err != nil {
			utils.JSON(w, http.StatusInternalServerError, false, nil, err.Error())
			return
		}

		utils.JSON(w, http.StatusOK, true, nil, "User deleted successfully")
	}
}
