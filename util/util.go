package util

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ReturnTableList(c *gin.Context, obj interface{}, table string) {
	c.JSON(http.StatusOK, gin.H{
		table: obj,
	})
}

func ReturnTableSingel(c *gin.Context, obj interface{}) {
	c.JSON(http.StatusOK, obj)
}

func ReturnErrorGeneral(c *gin.Context, msg interface{}, cause interface{}) {
	c.JSON(422, gin.H{
		"message": msg,
		"cause":   cause,
	})
}
