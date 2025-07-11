package controllers

import (
	"golang-social-media/config"
	"golang-social-media/models"
	"golang-social-media/utils"
	"net/http"

	"github.com/google/uuid"
	"github.com/gin-gonic/gin"
)

func RegisterUser(c *gin.Context) {
	var user models.User
	var existingUser models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		utils.RespondJSON(c, http.StatusBadRequest, "Invalid input", nil, err.Error())
		return
	}

	if err := config.DB.Where("email = ?", user.Email).First(&existingUser).Error; err == nil {
		utils.RespondJSON(c, http.StatusConflict, "Email already exists", nil, "User with this email already exists")
		return
	}

	user.ID = uuid.New()

	if err := config.DB.Create(&user).Error; err != nil {
		utils.RespondJSON(c, http.StatusInternalServerError, "Failed to register user", nil, err.Error())
		return
	}

	utils.RespondJSON(c, http.StatusCreated, "User registered", user, nil)
}

func GetUsers(c *gin.Context) {
	var users []models.User
	if err := config.DB.Find(&users).Error; err != nil {
		utils.RespondJSON(c, http.StatusInternalServerError, "Failed to retrieve users", nil, err.Error())
		return
	}

	utils.RespondJSON(c, http.StatusOK, "Users retrieved", users, nil)
}

func GetUserByID(c *gin.Context) {
	id := c.Param("id")

	if _, err := uuid.Parse(id); err != nil {
		utils.RespondJSON(c, http.StatusBadRequest, "Invalid UUID", nil, err.Error())
		return
	}

	var user models.User
	if err := config.DB.First(&user, "id = ?", id).Error; err != nil {
		utils.RespondJSON(c, http.StatusNotFound, "User not found", nil, err.Error())
		return
	}

	utils.RespondJSON(c, http.StatusOK, "User retrieved", user, nil)
}

func UpdateUser(c *gin.Context) {
	id := c.Param("id")

	if _, err := uuid.Parse(id); err != nil {
		utils.RespondJSON(c, http.StatusBadRequest, "Invalid UUID", nil, err.Error())
		return
	}

	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		utils.RespondJSON(c, http.StatusBadRequest, "Invalid input", nil, err.Error())
		return
	}

	if err := config.DB.Model(&models.User{}).Where("id = ?", id).Updates(user).Error; err != nil {
		utils.RespondJSON(c, http.StatusInternalServerError, "Failed to update user", nil, err.Error())
		return
	}

	utils.RespondJSON(c, http.StatusOK, "User updated", user, nil)
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")

	if _, err := uuid.Parse(id); err != nil {
		utils.RespondJSON(c, http.StatusBadRequest, "Invalid UUID", nil, err.Error())
		return
	}

	if err := config.DB.Where("id = ?", id).Delete(&models.User{}).Error; err != nil {
		utils.RespondJSON(c, http.StatusInternalServerError, "Failed to delete user", nil, err.Error())
		return
	}

	utils.RespondJSON(c, http.StatusOK, "User deleted", nil, nil)
}
