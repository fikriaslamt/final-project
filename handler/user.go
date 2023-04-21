package handler

import (
	"final-project/helper"
	"final-project/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h HttpServer) Register(c *gin.Context) {

	user := model.User{}

	err := c.BindJSON(&user)
	if err != nil {
		helper.BadRequest(c, err.Error())
		return
	}

	err = user.Validation()
	if err != nil {
		helper.BadRequest(c, err.Error())
		return
	}

	err = h.app.Validate(user)
	if err != nil {
		helper.BadRequest(c, err.Error())
		return
	}

	res, err := h.app.Register(user)
	if err != nil {
		helper.InternalServerError(c, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Register Success",
		"data": gin.H{
			"id":       res.ID,
			"email":    res.Email,
			"username": res.UserName,
		},
	})

}
func (h HttpServer) Login(c *gin.Context) {
	user := model.LoginCredentials{}

	err := c.BindJSON(&user)
	if err != nil {
		helper.BadRequest(c, err.Error())
		return
	}

	err = user.Validation()
	if err != nil {
		helper.BadRequest(c, err.Error())
		return
	}

	token, err := h.app.Login(user)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
