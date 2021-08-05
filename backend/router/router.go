package router

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"todo/api"
	_ "todo/docs"
	"todo/router/middleware"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.LoggerToFile(), middleware.Cors())
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	userRouter := r.Group("/api/user")
	{
		userRouter.POST("register", api.Register)
		userRouter.POST("login", api.Login)
	}
	taskRouter := r.Group("/api/tasks").Use(middleware.JWTAuth())
	{
		taskRouter.GET("", api.GetTasks)
		taskRouter.PUT("", api.CreateTask)
		taskRouter.POST("", api.UpdateTask)
		taskRouter.DELETE("", api.DeleteTask)
	}

	return r
}
