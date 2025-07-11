package utils

import (
	"github.com/gin-gonic/gin"
)

const (
	StatusOK             = 200
	StatusCreated        = 201
	StatusBadRequest     = 400
	StatusUnauthorized   = 401
	StatusForbidden      = 403
	StatusNotFound       = 404
	StatusInternalServer = 500
)

func RespondJSON(c *gin.Context, status int, message string, data interface{}, err interface{}) {
	c.JSON(status, gin.H{
		"message": message,
		"data":    data,
		"error":   err,
	})
}
