package controllers

import (
	"golang-social-media/config"
	"golang-social-media/models"
	"golang-social-media/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateComment(c *gin.Context) {
	var comment models.Comment

	if err := c.ShouldBindJSON(&comment); err != nil {
		utils.RespondJSON(c, http.StatusBadRequest, "Invalid input", nil, err.Error())
		return
	}

	if _, err := uuid.Parse(comment.UserID.String()); err != nil {
		utils.RespondJSON(c, http.StatusBadRequest, "Invalid user_id", nil, err.Error())
		return
	}

	if _, err := uuid.Parse(comment.PostID.String()); err != nil {
		utils.RespondJSON(c, http.StatusBadRequest, "Invalid post_id", nil, err.Error())
		return
	}

	comment.ID = uuid.New()

	if err := config.DB.Create(&comment).Error; err != nil {
		utils.RespondJSON(c, http.StatusInternalServerError, "Failed to create comment", nil, err.Error())
		return
	}

	utils.RespondJSON(c, http.StatusCreated, "Comment created", comment, nil)
}

func GetComments(c *gin.Context) {
	var comments []models.Comment
	if err := config.DB.Find(&comments).Error; err != nil {
		utils.RespondJSON(c, http.StatusInternalServerError, "Failed to retrieve comments", nil, err.Error())
		return
	}

	utils.RespondJSON(c, http.StatusOK, "Comments retrieved", comments, nil)
}

func GetCommentsByPostID(c *gin.Context) {
	id := c.Param("id")

	if _, err := uuid.Parse(id); err != nil {
		utils.RespondJSON(c, http.StatusBadRequest, "Invalid UUID", nil, err.Error())
		return
	}

	var comments []models.Comment
	if err := config.DB.Where("post_id = ?", id).Find(&comments).Error; err != nil {
		utils.RespondJSON(c, http.StatusInternalServerError, "Failed to retrieve comments", nil, err.Error())
		return
	}

	utils.RespondJSON(c, http.StatusOK, "Comments retrieved", comments, nil)
}

func DeleteComment(c *gin.Context) {
	id := c.Param("id")

	if _, err := uuid.Parse(id); err != nil {
		utils.RespondJSON(c, http.StatusBadRequest, "Invalid UUID", nil, err.Error())
		return
	}

	if err := config.DB.Delete(&models.Comment{}, "id = ?", id).Error; err != nil {
		utils.RespondJSON(c, http.StatusInternalServerError, "Failed to delete comment", nil, err.Error())
		return
	}

	utils.RespondJSON(c, http.StatusOK, "Comment deleted", nil, nil)
}
