package handler

import (
	"final-project/helper"
	"final-project/model"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func (h HttpServer) GetAllPhoto(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := userData["id"].(float64)
	res, err := h.app.GetAllPhoto(int(userID))
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

func (h HttpServer) GetOnePhoto(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := userData["id"].(float64)

	id := c.Param("id")
	photoID, err := strconv.Atoi(id)
	if err != nil {
		helper.BadRequest(c, err.Error())
		return
	}

	res, err := h.app.GetOnePhoto(int(userID), photoID)
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

func (h HttpServer) CreatePhoto(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := userData["id"].(float64)

	photo := model.Photo{}
	err := c.BindJSON(&photo)
	if err != nil {
		helper.BadRequest(c, err.Error())
		return
	}

	err = photo.Validation()
	if err != nil {
		helper.BadRequest(c, err.Error())
		return
	}

	res, err := h.app.CreatePhoto(int(userID), photo)
	if err != nil {
		helper.InternalServerError(c, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Create Photo Success",
		"data":    res,
	})

}

func (h HttpServer) UpdatePhoto(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := userData["id"].(float64)

	id := c.Param("id")
	photoID, err := strconv.Atoi(id)
	if err != nil {
		helper.BadRequest(c, err.Error())
		return
	}

	photo := model.Photo{}
	err = c.BindJSON(&photo)
	if err != nil {
		helper.BadRequest(c, err.Error())
		return
	}

	err = photo.Validation()
	if err != nil {
		helper.BadRequest(c, err.Error())
		return
	}

	res, err := h.app.UpdatePhoto(int(userID), photoID, photo)
	if err != nil {
		helper.InternalServerError(c, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Update Photo Success",
		"data":    res,
	})
}

func (h HttpServer) DeletePhoto(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := userData["id"].(float64)

	id := c.Param("id")
	photoID, err := strconv.Atoi(id)
	if err != nil {
		helper.BadRequest(c, err.Error())
		return
	}

	err = h.app.DeletePhoto(int(userID), photoID)
	if err != nil {
		helper.InternalServerError(c, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Delete Photo Success",
	})
}
