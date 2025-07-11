package controllers

import (
	"golang-social-media/config"
	"golang-social-media/models"
	"golang-social-media/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreatePost(c *gin.Context) {
	var post models.Post
	var validUser models.User

	if err := c.ShouldBindJSON(&post); err != nil {
		utils.RespondJSON(c, http.StatusBadRequest, "Invalid input", nil, err.Error())
		return
	}

	if _, err := uuid.Parse(post.UserID.String()); err != nil {
		utils.RespondJSON(c, http.StatusBadRequest, "Invalid user_id", nil, err.Error())
		return
	}

	if err := config.DB.First(&validUser, "id = ?", post.UserID).Error; err != nil {
		utils.RespondJSON(c, http.StatusNotFound, "User not found", nil, err.Error())
		return
	}

	post.ID = uuid.New()
	post.Created = time.Now()

	if err := config.DB.Create(&post).Error; err != nil {
		utils.RespondJSON(c, http.StatusInternalServerError, "Failed to create post", nil, err.Error())
		return
	}

	utils.RespondJSON(c, http.StatusCreated, "Post created", post, nil)
}

func GetPosts(c *gin.Context) {
	var posts []models.Post
	if err := config.DB.Find(&posts).Error; err != nil {
		utils.RespondJSON(c, http.StatusInternalServerError, "Failed to retrieve posts", nil, err.Error())
		return
	}

	utils.RespondJSON(c, http.StatusOK, "Posts retrieved", posts, nil)
}

func GetPostByID(c *gin.Context) {
	id := c.Param("id")

	if _, err := uuid.Parse(id); err != nil {
		utils.RespondJSON(c, http.StatusBadRequest, "Invalid UUID", nil, err.Error())
		return
	}

	var post models.Post
	if err := config.DB.First(&post, "id = ?", id).Error; err != nil {
		utils.RespondJSON(c, http.StatusNotFound, "Post not found", nil, err.Error())
		return
	}

	utils.RespondJSON(c, http.StatusOK, "Post retrieved", post, nil)
}

func GetPostsByUserID(c *gin.Context) {
	userID := c.Param("id")

	if _, err := uuid.Parse(userID); err != nil {
		utils.RespondJSON(c, http.StatusBadRequest, "Invalid user_id", nil, err.Error())
		return
	}

	var posts []models.Post
	if err := config.DB.Where("user_id = ?", userID).Find(&posts).Error; err != nil {
		utils.RespondJSON(c, http.StatusInternalServerError, "Failed to retrieve posts", nil, err.Error())
		return
	}

	utils.RespondJSON(c, http.StatusOK, "Posts retrieved", posts, nil)
}

func DeletePost(c *gin.Context) {
	id := c.Param("id")

	if _, err := uuid.Parse(id); err != nil {
		utils.RespondJSON(c, http.StatusBadRequest, "Invalid UUID", nil, err.Error())
		return
	}

	if err := config.DB.Delete(&models.Post{}, "id = ?", id).Error; err != nil {
		utils.RespondJSON(c, http.StatusInternalServerError, "Failed to delete post", nil, err.Error())
		return
	}

	utils.RespondJSON(c, http.StatusOK, "Post deleted", nil, nil)
}
