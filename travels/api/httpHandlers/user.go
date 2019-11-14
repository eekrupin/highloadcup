package httpHandlers

import (
	"github.com/gin-gonic/gin"
	"highloadcup/travels/db"
	"highloadcup/travels/models"
	"highloadcup/travels/modules"
	"log"
	"strconv"
	"time"
)

// /user
func User(c *gin.Context) {
	EntityResponse(c, "user")
}

func CreateUser(c *gin.Context) {

	var userRaw models.UserRaw
	err := c.BindJSON(userRaw)
	if err != nil || userRaw.Id == 0 {
		if err != nil {
			log.Println(err)
		}
		c.AbortWithStatus(404)
		return
	}

	raw, err := FirstFromTable("user", userRaw.Id)
	if err != nil || raw != nil {
		log.Println(err)
		c.AbortWithStatus(400)
		return
	}

	birth_date := time.Unix(int64(userRaw.Birth_date), 0)

	Age, _ := modules.MonthYearDiff(birth_date, time.Now())
	user := models.User{Id: userRaw.Id, Birth_date: birth_date, Email: userRaw.Email, Gender: userRaw.Gender, Last_name: userRaw.Last_name, Age: Age}

	tnx := db.DB.Txn(true)
	err = tnx.Insert("user", &user)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(404)
		return
	}
	tnx.Commit()

	resp := map[string]string{}
	c.JSON(200, resp)

	c.Abort()
}

func PostUser(c *gin.Context) {
	key := c.Param("id")
	if key == "new" {
		CreateUser(c)
		return
	}

	id, err := strconv.Atoi(key)
	if err != nil {
		c.AbortWithStatus(404)
		return
	}

	var userRaw models.UserRaw
	err = c.BindJSON(&userRaw)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(404)
		return
	}

	birth_date := time.Unix(int64(userRaw.Birth_date), 0)

	Age, _ := modules.MonthYearDiff(birth_date, time.Now())
	user := models.User{Id: uint32(id), Birth_date: birth_date, Email: userRaw.Email, Gender: userRaw.Gender, Last_name: userRaw.Last_name, Age: Age}

	tnx := db.DB.Txn(false)
	raw, err := tnx.First("user", "id", uint32(id))
	if err != nil || raw == nil {
		if err != nil {
			log.Println(err)
		}
		c.AbortWithStatus(404)
		return
	}

	tnx = db.DB.Txn(true)
	err = tnx.Insert("user", &user)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(404)
		return
	}
	tnx.Commit()

	resp := map[string]string{}
	c.JSON(200, resp)

	c.Abort()
}
