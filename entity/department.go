package entity

import (
	_connection "go-connect-mysql/connection"
	_util "go-connect-mysql/util"
	"strconv"

	"github.com/gin-gonic/gin"
)

type departments struct {
	Id          uint   `gorm:"primary_key, AUTO_INCREMENT"`
	Code        string `json:”code”`
	Description string `json:”description”`
	Name        string `json:”name”`
}

// get all data
func GetDepartments(c *gin.Context) {
	db, err := _connection.ConnectDatabase()
	if err == nil {
		allDep := []departments{}
		db.Find(&allDep)
		_util.ReturnTableList(c, allDep, "departments")
	} else {
		_util.ReturnErrorGeneral(c, "Cant get data departments", err)
	}
}

// get data by id
func GetDepartmentsById(id string, c *gin.Context) {
	i, err := strconv.Atoi(id)
	if err != nil {
		_util.ReturnErrorGeneral(c, "Cant get department data with id "+id, err)
	} else {
		db, err := _connection.ConnectDatabase()
		if err == nil {
			allDep := departments{}
			db.Find(&allDep, i)
			_util.ReturnTableSingel(c, allDep)
		} else {
			_util.ReturnErrorGeneral(c, "Cant update department data with id "+id, err)
		}
	}
}

// save data
func PostDepartments(c *gin.Context) {
	dep := departments{}

	if err := c.BindJSON(&dep); err != nil {
		_util.ReturnErrorGeneral(c, "Cant saving departments, please contact administrator.", err)
	}

	db, err := _connection.ConnectDatabase()
	if err == nil {
		db.Select("Id", "Code", "Name", "Description").Create(&dep)
		_util.ReturnTableSingel(c, dep)
	} else {
		_util.ReturnErrorGeneral(c, "Cant saving department, please contact administrator.", err)
	}
}

// get data by id
func PatchDepartments(id string, c *gin.Context) {
	i, err := strconv.Atoi(id)
	if err != nil {
		_util.ReturnErrorGeneral(c, "Cant update department data with id "+id, err)
	} else {
		db, err := _connection.ConnectDatabase()
		if err == nil {
			dep := departments{}
			if er := c.BindJSON(&dep); er != nil {
				_util.ReturnErrorGeneral(c, "Cant update departments, please contact administrator.", er)
			}
			result := db.Model(departments{}).Where("Id = ?", i).Updates(dep)
			if result.Error != nil {
				_util.ReturnErrorGeneral(c, "Cant update departments, please contact administrator.", result.Error)
			}
			_util.ReturnTableSingel(c, db.Find(&dep, i).Value)
		} else {
			_util.ReturnErrorGeneral(c, "Cant update department data with id "+id, err)
		}
	}
}
