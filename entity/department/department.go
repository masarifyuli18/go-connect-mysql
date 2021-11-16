package entity

import (
	_connection "go-connect-mysql/connection"
	_util "go-connect-mysql/util"
	"strconv"

	"github.com/gin-gonic/gin"
)

type department struct {
	Id          int
	Code        string
	Description string
	Name        string
}

func DepartmentController(Router *gin.Engine) {
	Router.GET("/departments", getDepartments)
	Router.GET("/departments/:id", func(c *gin.Context) {
		getDepartmentsById(c.Param("id"), c)
	})
}

// get all data
func getDepartments(c *gin.Context) {
	var department []department
	_, err := _connection.Db.Select(&department, "select * from department")

	if err == nil {
		_util.ReturnOk(c, department)
	} else {
		_util.ReturnErrorGeneral(c, "Cant get data departments", err)
	}
}

// get data by id
func getDepartmentsById(id string, c *gin.Context) {
	i, err := strconv.Atoi(id)
	if err != nil {
		_util.ReturnErrorGeneral(c, "Cant get department data with id "+id, err)
	} else {
		var department department
		err := _connection.Db.SelectOne(&department, "select * from department where id =? LIMIT 1", i)

		if err == nil {
			_util.ReturnOk(c, department)
		} else {
			_util.ReturnErrorGeneral(c, "Cant get department data with id "+id, err)
		}
	}
}
