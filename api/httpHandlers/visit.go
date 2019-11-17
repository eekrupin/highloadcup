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

// /visit
func Visit(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		log.Println(err)
		c.AbortWithStatus(404)
		return
	}

	rec, err := db.RDB.FindByPrimaryKeyFrom(models.VisitTable, uint32(id))
	if err != nil || rec == nil {
		if err != nil {
			log.Println(err)
		}
		c.AbortWithStatus(404)
		return
	}

	resp, err := rec.(*models.Visit).MarshalJSON()
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(404)
		return
	}

	c.Data(200, modules.JSONContentType, resp)

	c.Abort()
}

func CreateVisit(c *gin.Context) {

	var visitRaw models.VisitRaw
	err := c.BindJSON(&visitRaw)
	if err != nil || visitRaw.Id == 0 {
		if err != nil {
			log.Println(err)
		}
		c.AbortWithStatus(404)
		return
	}

	rec, err := db.RDB.FindByPrimaryKeyFrom(models.VisitTable, uint32(visitRaw.Id))
	if rec != nil {
		log.Println(err)
		c.AbortWithStatus(400)
		return
	}

	visited_at := time.Unix(int64(visitRaw.Visited_at), 0)

	visit := models.Visit{Id: visitRaw.Id, Location: visitRaw.Location, User: visitRaw.User, Visited_at: visited_at, Mark: visitRaw.Mark}

	err = db.RDB.Save(&visit)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(404)
		return
	}

	resp := map[string]string{}
	c.JSON(200, resp)

	c.Abort()
}

func PostVisit(c *gin.Context) {
	key := c.Param("id")
	if key == "new" {
		CreateVisit(c)
		return
	}

	id, err := strconv.Atoi(key)
	if err != nil {
		c.AbortWithStatus(404)
		return
	}

	var visitRaw models.VisitRaw
	err = c.BindJSON(&visitRaw)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(404)
		return
	}

	visited_at := time.Unix(int64(visitRaw.Visited_at), 0)

	visit := models.Visit{Id: uint32(id), Location: visitRaw.Location, User: visitRaw.User, Visited_at: visited_at, Mark: visitRaw.Mark}

	rec, err := db.RDB.FindByPrimaryKeyFrom(models.VisitTable, uint32(id))
	if err != nil || rec == nil {
		if err != nil {
			log.Println(err)
		}
		c.AbortWithStatus(404)
		return
	}

	err = db.RDB.Save(&visit)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(404)
		return
	}

	resp := map[string]string{}
	c.JSON(200, resp)

	c.Abort()
}
