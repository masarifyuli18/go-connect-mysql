package connection

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/go-gorp/gorp"
	_ "github.com/go-sql-driver/mysql"
)

var Db = connectDatabase()

func connectDatabase() *gorp.DbMap {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/simpleapp")
	checkErr(err, "sql.Open failed")
	dbmap := &gorp.DbMap{
		Db:              db,
		Dialect:         gorp.MySQLDialect{"InnoDB", "UTF8"},
		TypeConverter:   nil,
		ExpandSliceArgs: false,
	}
	err = dbmap.CreateTablesIfNotExists()
	checkErr(err, "Create tables failed")
	return dbmap
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
