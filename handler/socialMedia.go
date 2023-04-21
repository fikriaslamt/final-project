package handler

import (
	"final-project/helper"
	"final-project/model"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func (h HttpServer) GetAllSocialMedia(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := userData["id"].(float64)
	res, err := h.app.GetAllSocialMedia(int(userID))
	if err != nil {
		helper.InternalServerError(c, err.Error())
		return
	}
	if len(res) == 0 {
		helper.NoContent(c)
		return
	}

	helper.Ok(c, res)
}

func (h HttpServer) GetOneSocialMedia(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := userData["id"].(float64)

	id := c.Param("id")
	sosmedID, err := strconv.Atoi(id)
	if err != nil {
		helper.BadRequest(c, err.Error())
		return
	}

	res, err := h.app.GetOneSocialMedia(int(userID), sosmedID)
	if err != nil {
		if err.Error() == helper.ErrNotFound {
			helper.NotFound(c, err.Error())
			return
		}
		helper.InternalServerError(c, err.Error())
		return
	}

	helper.Ok(c, res)
}

func (h HttpServer) CreateSocialMedia(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := userData["id"].(float64)

	sosmed := model.SocialMedia{}
	err := c.BindJSON(&sosmed)
	if err != nil {
		helper.BadRequest(c, err.Error())
		return
	}

	err = sosmed.Validation()
	if err != nil {
		helper.BadRequest(c, err.Error())
		return
	}

	res, err := h.app.CreateSocialMedia(int(userID), sosmed)
	if err != nil {
		helper.BadRequest(c, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Create Social Media Success",
		"data":    res,
	})
}

func (h HttpServer) UpdateSocialMedia(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := userData["id"].(float64)

	id := c.Param("id")
	sosmedID, err := strconv.Atoi(id)
	if err != nil {
		helper.BadRequest(c, err.Error())
		return
	}

	sosmed := model.SocialMedia{}
	err = c.BindJSON(&sosmed)
	if err != nil {
		helper.BadRequest(c, err.Error())
		return
	}

	err = sosmed.Validation()
	if err != nil {
		helper.BadRequest(c, err.Error())
		return
	}

	res, err := h.app.UpdateSocialMedia(int(userID), sosmedID, sosmed)
	if err != nil {
		helper.InternalServerError(c, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Update Social Media Success",
		"data":    res,
	})
}

func (h HttpServer) DeleteSocialMedia(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := userData["id"].(float64)

	id := c.Param("id")
	sosmedID, err := strconv.Atoi(id)
	if err != nil {
		helper.BadRequest(c, err.Error())
		return
	}

	err = h.app.DeleteSocialMedia(int(userID), sosmedID)
	if err != nil {
		helper.InternalServerError(c, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Delete Social Media Success",
	})
}
