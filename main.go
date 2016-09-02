package main

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/go-sql-driver/mysql"
	"log"
	c "github.com/mrtomyum/gin/controller"
	"github.com/gin-gonic/gin"
	_ "github.com/denisenkom/go-mssqldb"
)



const (
	DB_HOST = "tcp(nava.work:3306)"
	DB_NAME = "stock"
	DB_USER = "root"
	DB_PASS = "mypass"
)

func NewDB(driver, dsn string) (*sqlx.DB, error){
	db := sqlx.MustConnect(driver, dsn)
	return db, nil
}
func SetupRouter(e *c.API) *gin.Engine{
	gin.DisableBindValidation()
	app := gin.New()
	app.GET("/", e.Index)
	app.GET("/users", e.GetAllUser)
	app.POST("/users", e.NewUser)
	return app
}

func main() {
	var myDSN = DB_USER + ":" + DB_PASS + "@" + DB_HOST + "/" + DB_NAME + "?parseTime=true"
	myDB, err := NewDB("mysql", myDSN)
	if err != nil {
		log.Panic("NewDB() Error:", err)
	}
	defer myDB.Close()

	e := &c.API{DB: myDB}
	app := SetupRouter(e)
	app.Run(":8080")
}