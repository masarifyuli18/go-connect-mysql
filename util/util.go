package util

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ReturnOk(c *gin.Context, obj interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"obj": obj,
	})
}

func ReturnErrorGeneral(c *gin.Context, msg interface{}, cause interface{}) {
	c.JSON(422, gin.H{
		"message": msg,
		"cause":   cause,
	})
}
