package config

import (
	_connection "go-connect-mysql/connection"
	_department "go-connect-mysql/entity/department"

	"github.com/gin-gonic/gin"
	ginserver "github.com/go-oauth2/gin-server"
)

var Router *gin.Engine

func Controller() {
	Router = gin.Default()
	Router.Use(_connection.Cors())

	// oauth controller
	auth := Router.Group("/oauth2")
	{
		auth.GET("/token", ginserver.HandleTokenRequest)
	}

	_department.DepartmentController(Router)
	Router.Run()
}
