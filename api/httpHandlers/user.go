package httpHandlers

import (
	"github.com/eekrupin/hlc-travels/db"
	"github.com/eekrupin/hlc-travels/models"
	"github.com/eekrupin/hlc-travels/modules"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
	"time"
)

// /user
func User(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		log.Println(err)
		c.AbortWithStatus(404)
		return
	}

	rec, err := db.RDB.FindByPrimaryKeyFrom(models.UserTable, uint32(id))
	if err != nil || rec == nil {
		if err != nil {
			log.Println(err)
		}
		c.AbortWithStatus(404)
		return
	}

	user := rec.(*models.User)
	//user.Age, _ = modules.MonthYearDiff(user.Birth_date, time.Now())
	resp, err := user.MarshalJSON()
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(404)
		return
	}

	c.Data(200, modules.JSONContentType, resp)

	c.Abort()
}

func CreateUser(c *gin.Context) {

	var userRaw models.UserRaw
	err := c.BindJSON(&userRaw)
	if err != nil || userRaw.Id == 0 {
		if err != nil {
			log.Println(err)
		}
		c.AbortWithStatus(404)
		return
	}

	rec, err := db.RDB.FindByPrimaryKeyFrom(models.UserTable, uint32(userRaw.Id))
	if rec != nil {
		log.Println(err)
		c.AbortWithStatus(400)
		return
	}
	user := models.UserFromRaw(&userRaw)

	err = db.RDB.Save(user)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(404)
		return
	}

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

	//Age, _ := modules.MonthYearDiff(birth_date, time.Now())
	user := models.User{Id: uint32(id), Birth_date: birth_date, Email: userRaw.Email, Gender: userRaw.Gender, Last_name: userRaw.Last_name, First_name: userRaw.First_name}

	rec, err := db.RDB.FindByPrimaryKeyFrom(models.UserTable, uint32(id))
	if err != nil || rec == nil {
		if err != nil {
			log.Println(err)
		}
		c.AbortWithStatus(404)
		return
	}

	err = db.RDB.Save(&user)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(404)
		return
	}

	resp := map[string]string{}
	c.JSON(200, resp)

	c.Abort()
}
