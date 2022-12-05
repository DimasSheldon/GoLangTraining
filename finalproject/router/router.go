package router

import (
	"finalproject/controller"
	"finalproject/middleware"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", controller.UserRegister)
		userRouter.POST("/login", controller.UserLogin)
		userRouter.PUT("/", middleware.Authentication(), controller.UpdateUser)
		userRouter.DELETE("/", middleware.Authentication(), controller.DeleteUser)
	}

	photoRouter := r.Group("/photos")
	{
		photoRouter.Use(middleware.Authentication())
		photoRouter.POST("/", controller.CreatePhoto)
		photoRouter.GET("/", controller.GetPhoto)
		photoRouter.PUT("/:photoId", middleware.PhotoAuthorization(), controller.UpdatePhoto)
		photoRouter.DELETE("/:photoId", middleware.PhotoAuthorization(), controller.DeletePhoto)
	}

	commentRouter := r.Group("/comments")
	{
		commentRouter.Use(middleware.Authentication())
		commentRouter.POST("/", controller.CreateComment)
		commentRouter.GET("/", controller.GetComment)
		commentRouter.PUT("/:commentId", middleware.CommentAuthorization(), controller.UpdateComment)
		commentRouter.DELETE("/:commentId", middleware.CommentAuthorization(), controller.DeleteComment)
	}

	socialMediaRouter := r.Group("/socialmedias")
	{
		socialMediaRouter.Use(middleware.Authentication())
		socialMediaRouter.POST("/", controller.CreateSocialMedia)
		socialMediaRouter.GET("/", controller.GetSocialMedia)
		socialMediaRouter.PUT("/:socialMediaId", middleware.SocialMediaAuthorization(), controller.UpdateSocialMedia)
		socialMediaRouter.DELETE("/:socialMediaId", middleware.SocialMediaAuthorization(), controller.DeleteSocialMedia)
	}

	return r
}
