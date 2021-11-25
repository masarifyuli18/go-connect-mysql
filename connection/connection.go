package connection

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func ConnectDatabase() (*gorm.DB, error) {
	// dsn := "sqlserver://root:root@localhost:3306?database=simpleapp?net_write_timeout=6000"
	db, err := gorm.Open("mysql", "root:root@(localhost:3306)/simpleapp?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	// defer db.Close()
	return db, nil
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}
