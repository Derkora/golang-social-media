package controllers

import (
	"golang-social-media/config"
	"golang-social-media/models"
	"golang-social-media/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func LikePost(c *gin.Context) {
	var like models.Like
	var existingLike models.Like

	if err := c.ShouldBindJSON(&like); err != nil {
		utils.RespondJSON(c, http.StatusBadRequest, "Invalid input", nil, err.Error())
		return
	}

	if _, err := uuid.Parse(like.UserID.String()); err != nil {
		utils.RespondJSON(c, http.StatusBadRequest, "Invalid user_id", nil, err.Error())
		return
	}

	if _, err := uuid.Parse(like.PostID.String()); err != nil {
		utils.RespondJSON(c, http.StatusBadRequest, "Invalid post_id", nil, err.Error())
		return
	}

	if err := config.DB.Where("user_id = ? AND post_id = ?", like.UserID, like.PostID).First(&existingLike).Error; err == nil {
		utils.RespondJSON(c, http.StatusConflict, "Like already exists", nil, "User has already liked this post")
		return
	}

	like.ID = uuid.New()

	if err := config.DB.Create(&like).Error; err != nil {
		utils.RespondJSON(c, http.StatusInternalServerError, "Failed to create like", nil, err.Error())
		return
	}

	utils.RespondJSON(c, http.StatusCreated, "Like created", like, nil)
}

func GetLikes(c *gin.Context) {
	var likes []models.Like
	if err := config.DB.Find(&likes).Error; err != nil {
		utils.RespondJSON(c, http.StatusInternalServerError, "Failed to retrieve likes", nil, err.Error())
		return
	}

	utils.RespondJSON(c, http.StatusOK, "Likes retrieved", likes, nil)
}

func GetLikesByPostID(c *gin.Context) {
	id := c.Param("id")

	if _, err := uuid.Parse(id); err != nil {
		utils.RespondJSON(c, http.StatusBadRequest, "Invalid UUID", nil, err.Error())
		return
	}

	var likes []models.Like
	if err := config.DB.Where("post_id = ?", id).Find(&likes).Error; err != nil {
		utils.RespondJSON(c, http.StatusNotFound, "Likes not found for post", nil, err.Error())
		return
	}

	utils.RespondJSON(c, http.StatusOK, "Likes retrieved for post", likes, nil)
}

func GetLikesByUserID(c *gin.Context) {
	id := c.Param("id")

	if _, err := uuid.Parse(id); err != nil {
		utils.RespondJSON(c, http.StatusBadRequest, "Invalid UUID", nil, err.Error())
		return
	}

	var likes []models.Like
	if err := config.DB.Where("user_id = ?", id).Find(&likes).Error; err != nil {
		utils.RespondJSON(c, http.StatusNotFound, "Likes not found for user", nil, err.Error())
		return
	}

	utils.RespondJSON(c, http.StatusOK, "Likes retrieved for user", likes, nil)
}

func UnlikePost(c *gin.Context) {
	id := c.Param("id")

	if _, err := uuid.Parse(id); err != nil {
		utils.RespondJSON(c, http.StatusBadRequest, "Invalid UUID", nil, err.Error())
		return
	}

	var like models.Like
	if err := config.DB.First(&like, "id = ?", id).Error; err != nil {
		utils.RespondJSON(c, http.StatusNotFound, "Like not found", nil, err.Error())
		return
	}

	if err := config.DB.Delete(&like).Error; err != nil {
		utils.RespondJSON(c, http.StatusInternalServerError, "Failed to delete like", nil, err.Error())
		return
	}

	utils.RespondJSON(c, http.StatusOK, "Like deleted", nil, nil)
}
