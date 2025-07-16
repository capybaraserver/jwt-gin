package controllers

import (
	"jwt-gin/models"
	"jwt-gin/utils/ret"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusInternalServerError, ret.RetFailMsg("Internal Server Error"))
		return
	}

	user, err := models.GetUserByID(userID.(int64))
	if err != nil {
		c.JSON(http.StatusInternalServerError, ret.RetFailMsg("Internal Server Error"))
		return
	}

	c.JSON(http.StatusOK, ret.RetSuccessData(user))
}
