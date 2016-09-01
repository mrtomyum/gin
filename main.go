package main

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"encoding/json"
	"os"
	c "github.com/mrtomyum/gin/controller"
	"github.com/gin-gonic/gin"
)

type Config struct {
	DBHost string `json:"db_host"`
	DBName string `json:"db_name"`
	DBUser string `json:"db_user"`
	DBPass string `json:"db_pass"`
}

func loadConfig() *Config {
	file, _ := os.Open("config.json")
	decoder := json.NewDecoder(file)
	config := new(Config)
	err := decoder.Decode(&config)
	if err != nil {
		log.Println("error:", err)
	}
	return config
}

func NewDB(dsn string) (*sqlx.DB, error){
	db := sqlx.MustConnect("mysql", dsn)
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
	config := loadConfig()
	var dsn = config.DBUser + ":" + config.DBPass + "@" + config.DBHost + "/" + config.DBName + "?parseTime=true"
	db, err := NewDB(dsn)
	if err != nil {
		log.Panic("NewDB() Error:", err)
	}
	defer db.Close()
	e := &c.API{DB: db}
	app := SetupRouter(e)
	app.Run(":8080")
}