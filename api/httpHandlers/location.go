package httpHandlers

import (
	"github.com/eekrupin/hlc-travels/db"
	"github.com/eekrupin/hlc-travels/models"
	"github.com/eekrupin/hlc-travels/modules"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

// /location
func Location(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		log.Println(err)
		c.AbortWithStatus(404)
		return
	}

	rec, err := db.RDB.FindByPrimaryKeyFrom(models.LocationTable, uint32(id))
	if err != nil || rec == nil {
		if err != nil {
			log.Println(err)
		}
		c.AbortWithStatus(404)
		return
	}

	resp, err := rec.(*models.Location).MarshalJSON()
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(404)
		return
	}

	c.Data(200, modules.JSONContentType, resp)

	c.Abort()
}

func CreateLocation(c *gin.Context) {

	var locationRaw models.Location
	err := c.BindJSON(&locationRaw)
	if err != nil || locationRaw.Id == 0 {
		if err != nil {
			log.Println(err)
		}
		c.AbortWithStatus(404)
		return
	}

	rec, err := db.RDB.FindByPrimaryKeyFrom(models.UserTable, uint32(locationRaw.Id))
	if rec != nil {
		log.Println(err)
		c.AbortWithStatus(400)
		return
	}

	location := models.Location{Id: locationRaw.Id, Place: locationRaw.Place, Country: locationRaw.Country, City: locationRaw.City, Distance: locationRaw.Distance}

	err = db.RDB.Save(&location)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(404)
		return
	}

	resp := map[string]string{}
	c.JSON(200, resp)

	c.Abort()
}

func PostLocation(c *gin.Context) {
	key := c.Param("id")
	if key == "new" {
		CreateLocation(c)
		return
	}

	id, err := strconv.Atoi(key)
	if err != nil {
		c.AbortWithStatus(404)
		return
	}

	var locationRaw models.Location
	err = c.BindJSON(&locationRaw)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(404)
		return
	}

	location := models.Location{Id: uint32(id), Place: locationRaw.Place, Country: locationRaw.Country, City: locationRaw.City, Distance: locationRaw.Distance}

	rec, err := db.RDB.FindByPrimaryKeyFrom(models.LocationTable, uint32(id))
	if err != nil || rec == nil {
		if err != nil {
			log.Println(err)
		}
		c.AbortWithStatus(404)
		return
	}

	err = db.RDB.Save(&location)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(404)
		return
	}

	resp := map[string]string{}
	c.JSON(200, resp)

	c.Abort()
}
