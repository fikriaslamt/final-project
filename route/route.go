package route

import (
	"final-project/handler"
	"final-project/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterApi(r *gin.Engine, handler handler.HttpServer) {
	user := r.Group("/users") // prefix
	{
		user.POST("/register", handler.Register)
		user.POST("/login", handler.Login)
	}

	comment := r.Group("/comments") // prefix
	{
		comment.Use(middleware.Authentication())
		comment.GET("/:photo_id", handler.GetAllComment)
		comment.GET("/:photo_id/:id", handler.GetOneComment)
		comment.POST("/create/:photo_id", handler.CreateComment)
		comment.PUT("/update/:id", middleware.CommentAuthorization(), handler.UpdateComment)
		comment.DELETE("/delete/:id", middleware.CommentAuthorization(), handler.DeleteComment)
	}

	photo := r.Group("/photos") // prefix
	{
		photo.Use(middleware.Authentication())
		photo.GET("", handler.GetAllPhoto)
		photo.GET("/:id", handler.GetOnePhoto)
		photo.POST("/create", handler.CreatePhoto)
		photo.PUT("/update/:id", middleware.PhotoAuthorization(), handler.UpdatePhoto)
		photo.DELETE("/delete/:id", middleware.PhotoAuthorization(), handler.DeletePhoto)
	}

	sosmed := r.Group("/social-media") // prefix
	{
		sosmed.Use(middleware.Authentication())
		sosmed.GET("", handler.GetAllSocialMedia)
		sosmed.GET("/:id", handler.GetOneSocialMedia)
		sosmed.POST("/create", handler.CreateSocialMedia)
		sosmed.PUT("/update/:id", middleware.SocialMediaAuthorization(), handler.UpdateSocialMedia)
		sosmed.DELETE("/delete/:id", middleware.SocialMediaAuthorization(), handler.DeleteSocialMedia)
	}
}
