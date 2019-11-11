package httpHandlers

import (
	"github.com/gin-gonic/gin"
	"highloadcup/travels/db"
	"highloadcup/travels/models"
	"strconv"
	"time"
)

// /user
func User(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.AbortWithStatus(404)
		return
	}

	tnx := db.DB.Txn(false)
	raw, err := tnx.First("user", "id", uint32(id))
	if err != nil || raw == nil {
		c.AbortWithStatus(404)
	}

	resp := map[string]string{"status": raw.(*models.User).Last_name}
	c.JSON(200, resp)
	//c.Set(config.KeyResponse, resp)
	//c.JSON(http.StatusOK, map[string]string{"error": err.Error()})

	c.Abort()
}

func CreateUser(c *gin.Context) {

	var userRaw models.UserRaw
	err := c.BindJSON(&userRaw)
	if err != nil || userRaw.Id == 0 {
		c.AbortWithStatus(404)
		return
	}

	tnx := db.DB.Txn(false)
	raw, err := tnx.First("user", "id", userRaw.Id)
	if err != nil || raw != nil {
		c.AbortWithStatus(400)
		return
	}

	birth_date := time.Unix(int64(userRaw.Birth_date), 0)

	Age, _ := monthYearDiff(birth_date, time.Now())
	user := models.User{Id: userRaw.Id, Birth_date: birth_date, Email: userRaw.Email, Gender: userRaw.Gender, Last_name: userRaw.Last_name, Age: Age}

	tnx = db.DB.Txn(true)
	err = tnx.Insert("user", &user)
	if err != nil {
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
		c.AbortWithStatus(404)
		return
	}

	birth_date := time.Unix(int64(userRaw.Birth_date), 0)

	Age, _ := monthYearDiff(birth_date, time.Now())
	user := models.User{Id: uint32(id), Birth_date: birth_date, Email: userRaw.Email, Gender: userRaw.Gender, Last_name: userRaw.Last_name, Age: Age}

	tnx := db.DB.Txn(false)
	raw, err := tnx.First("user", "id", uint32(id))
	if err != nil || raw == nil {
		c.AbortWithStatus(404)
		return
	}

	tnx = db.DB.Txn(true)
	err = tnx.Insert("user", &user)
	if err != nil {
		c.AbortWithStatus(404)
		return
	}
	tnx.Commit()

	resp := map[string]string{}
	c.JSON(200, resp)

	c.Abort()
}

func monthYearDiff(a, b time.Time) (years, months int) {
	m := a.Month()
	for a.Before(b) {
		a = a.Add(time.Hour * 24)
		m2 := a.Month()
		if m2 != m {
			months++
		}
		m = m2
	}
	years = months / 12
	months = months % 12
	return
}
