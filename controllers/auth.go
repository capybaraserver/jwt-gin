package controllers

import (
	"jwt-gin/models"
	"net/http"

	"jwt-gin/utils/ret"

	"github.com/gin-gonic/gin"
)

type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var input LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, ret.RetFailMsg(err.Error()))
		return
	}

	u := models.User{}

	u.Username = input.Username
	u.Password = input.Password

	token, err := models.LoginCheck(u.Username, u.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, ret.RetFailMsg("username or password is incorrect."))
		return
	}

	c.JSON(http.StatusOK, ret.RetSuccessData(token))
}

type RegisterInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Register(c *gin.Context) {
	var input RegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, ret.RetFailMsg(err.Error()))
		return
	}

	u := models.User{}
	u.Username = input.Username
	u.Password = input.Password
	_, err := u.SaveUser()
	if err != nil {
		c.JSON(http.StatusBadRequest, ret.RetFailMsg(err.Error()))
		return
	}

	c.JSON(http.StatusOK, ret.RetSuccess())
}
