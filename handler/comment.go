package handler

import (
	"final-project/helper"
	"final-project/model"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func (h HttpServer) GetAllComment(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := userData["id"].(float64)

	id := c.Param("photo_id")
	PhotoID, err := strconv.Atoi(id)
	if err != nil {
		helper.BadRequest(c, err.Error())
		return
	}

	res, err := h.app.GetAllComment(int(userID), PhotoID)
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

func (h HttpServer) GetOneComment(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := userData["id"].(float64)

	id := c.Param("id")
	CommentID, err := strconv.Atoi(id)
	if err != nil {
		helper.BadRequest(c, err.Error())
		return
	}

	res, err := h.app.GetOneComment(int(userID), CommentID)
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

func (h HttpServer) CreateComment(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := userData["id"].(float64)

	id := c.Param("photo_id")
	PhotoID, err := strconv.Atoi(id)
	if err != nil {
		helper.BadRequest(c, err.Error())
		return
	}

	Comment := model.Comment{}
	err = c.BindJSON(&Comment)
	if err != nil {
		helper.BadRequest(c, err.Error())
		return
	}

	err = Comment.Validation()
	if err != nil {
		helper.BadRequest(c, err.Error())
		return
	}

	res, err := h.app.CreateComment(int(userID), PhotoID, Comment)
	if err != nil {
		helper.InternalServerError(c, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Create Comment Success",
		"data":    res,
	})

}

func (h HttpServer) UpdateComment(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := userData["id"].(float64)

	id := c.Param("id")
	CommentID, err := strconv.Atoi(id)
	if err != nil {
		helper.BadRequest(c, err.Error())
		return
	}

	Comment := model.Comment{}
	err = c.BindJSON(&Comment)
	if err != nil {
		helper.BadRequest(c, err.Error())
		return
	}

	err = Comment.Validation()
	if err != nil {
		helper.BadRequest(c, err.Error())
		return
	}

	res, err := h.app.UpdateComment(int(userID), CommentID, Comment)
	if err != nil {
		helper.InternalServerError(c, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Update Comment Success",
		"data":    res,
	})
}

func (h HttpServer) DeleteComment(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := userData["id"].(float64)

	id := c.Param("id")
	CommentID, err := strconv.Atoi(id)
	if err != nil {
		helper.BadRequest(c, err.Error())
		return
	}

	err = h.app.DeleteComment(int(userID), CommentID)
	if err != nil {
		helper.InternalServerError(c, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Delete Comment Success",
	})
}
