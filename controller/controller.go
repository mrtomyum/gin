package controller

import (
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	m "github.com/mrtomyum/gin/model"
)

type API struct {
	DB *sqlx.DB
}



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

	// Todo อย่างที่คิด c.BindJSON ต้องมี JSON Field ที่ตรงกับ Struct เป๊ะ
	// ไม่งั้นจะไม่อ่าน ต่างจาก json.NewDecoder จะ Matching Field ที่ตรงกันให้เอง
	// ซึ่งสามารถป้อนเฉพาะฟิลด์ที่ต้องการได้
	if  c.BindJSON(&u) != nil {
		c.JSON(400, u)
	} else {
		users, err := u.New(a.DB)
		if err != nil {
			log.Println(err)
		}

		log.Println(users)
		c.JSON(http.StatusOK, users)
	}
	return
}
