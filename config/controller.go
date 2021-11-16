package config

import (
	_connection "go-connect-mysql/connection"
	_department "go-connect-mysql/entity/department"

	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func Controller() {
	Router = gin.Default()
	Router.Use(_connection.Cors())
	_department.DepartmentController(Router)
	Router.Run()
}
