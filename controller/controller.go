package controller

import (
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
	m "github.com/mrtomyum/gin/model"
)

func (a *API) Index(c *gin.Context) {
	content := gin.H{"Hello": "World"}
	c.JSON(http.StatusOK, content)
}

func (a *API) GetAllUser(c *gin.Context) {
	log.Println("GetAllUser()")
	c.Header("Content-Type", "application/json")
	c.Header("Access-Control-Allow-Origin", "*")
	var u m.User

	users, err := u.All(a.DB)
	if err != nil {
		log.Println(err)
	}
	log.Println(users)
	c.JSON(http.StatusOK, users)
	return
}

func (a *API) NewUser(c *gin.Context) {
	log.Println("NewUser()")
	body := c.Request.Body
	log.Println(body)
	c.Header("Content-Type", "application/json")
	c.Header("Access-Control-Allow-Origin", "*")
	var u m.User
	rs := Response{}
	if  err := c.BindJSON(&u); err != nil {
		rs.Status = ERROR
		rs.Message = err.Error()
		c.JSON(400, rs)
	} else {
		users, err := u.New(a.DB)
		if err != nil {
			log.Println(err)
		}
		log.Println(users)
		rs.Status = SUCCESS
		rs.Data = users
		c.JSON(http.StatusOK, rs)
	}
	return
}
