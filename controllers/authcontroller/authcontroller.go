package authcontroller

import (
	"encoding/json"
	"net/http"

	"github.com/roshif-study/task-5-vix-btpns--ROSHIF-/helper"
	"github.com/roshif-study/task-5-vix-btpns--ROSHIF-/models"
	"golang.org/x/crypto/bcrypt"
)

func Register(w http.ResponseWriter, r *http.Request) {

	var userInput models.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&userInput); err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
		}
	defer r.Body.Close()

	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(userInput.Password), bcrypt.DefaultCost)
	userInput.Password = string(hashPassword)

	if err := models.DB.Create(&userInput).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := map[string]string{"message": "Success"}
	helper.ResponseJSON(w, http.StatusOK, response)
}

func Login(w http.ResponseWriter, r *http.Request) {

}

func Logout(w http.ResponseWriter, r *http.Request) {

}