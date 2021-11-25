package main

import (
	"fmt"
	_config "go-connect-mysql/config"
	_entity "go-connect-mysql/entity"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("starting web server at http://localhost:8080/")

	server := gin.New()
	server.Use(gin.Logger())
	server.POST("/login", _config.Login)
	api := server.Group("/api", _config.TokenAuthMiddleware())
	{
		api.GET("/departments", _entity.GetDepartments)
		api.GET("/departments/:id", func(c *gin.Context) {
			_entity.GetDepartmentsById(c.Param("id"), c)
		})
		api.POST("/departments", _entity.PostDepartments)
		api.PATCH("/departments/:id", func(c *gin.Context) {
			_entity.PatchDepartments(c.Param("id"), c)
		})
	}

	port := "9090"
	server.Run(":" + port)
}
