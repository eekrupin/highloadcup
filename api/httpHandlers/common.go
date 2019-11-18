package httpHandlers

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

func RequestStringParam(c *gin.Context, param string) (request interface{}, ok bool) {
	values := c.Request.URL.Query()
	_, ok = values[param]
	str := values.Get(param)
	if ok && str == "" {
		c.Keys["err"] = true
		return nil, false
	} else {
		return str, ok
	}
}

func RequestIntParam(c *gin.Context, param string) (request interface{}, ok bool) {
	var err error
	values := c.Request.URL.Query()
	_, ok = values[param]
	str := values.Get(param)
	if ok {
		request, err = strconv.Atoi(str)
		if err != nil {
			c.Keys["err"] = true
			return nil, false
		}
		return request, ok
	} else {
		return nil, false
	}
}

func RequestAgeParamToTimeUnixViaCurrentTime(c *gin.Context, param string) (request interface{}, ok bool) {
	values := c.Request.URL.Query()
	_, ok = values[param]
	str := values.Get(param)
	if ok {
		year, err := strconv.Atoi(str)
		if err != nil {
			c.Keys["err"] = true
			return nil, false
		}
		request = time.Now().AddDate(-year, 0, 0).Unix()
		return request, ok
	} else {
		return nil, false
	}
}

func addConvArgs(filter *string, args *[]interface{}, c *gin.Context, request string, convert func(*gin.Context, string) (interface{}, bool), filterString string) {
	if fromDate, ok := convert(c, request); ok {
		*filter = *filter + " " + filterString
		*args = append(*args, fromDate)
	}
}

//
//import (
//	"github.com/gin-gonic/gin"
//	"highloadcup/travels/db"
//	"highloadcup/travels/models"
//	"highloadcup/travels/modules"
//	"log"
//	"strconv"
//	"time"
//)
//
//type EntityObject interface {
//	MarshalJSON() ([]byte, error)
//}
//
//func FirstFromTable(table string, id uint32) (raw interface{}, err error) {
//	rec, err := db.RDB.FindByPrimaryKeyFrom(models.UserTable, uint32(id))
//	if err != nil {
//		log.Fatal(err)
//	}
//	tnx := db.RDB.Txn(false)
//	raw, err = tnx.First(table, "id", uint32(id))
//	return
//}
//
//func EntityResponse(c *gin.Context, table string) {
//	id, err := strconv.Atoi(c.Param("id"))
//
//	if err != nil {
//		log.Println(err)
//		c.AbortWithStatus(404)
//		return
//	}
//
//	raw, err := FirstFromTable(table, uint32(id))
//	if err != nil || raw == nil {
//		if err != nil {
//			log.Println(err)
//		}
//		c.AbortWithStatus(404)
//		return
//	}
//
//	resp, err := raw.(EntityObject).MarshalJSON()
//	if err != nil {
//		log.Println(err)
//		c.AbortWithStatus(404)
//		return
//	}
//
//	c.Data(200, modules.JSONContentType, resp)
//
//	c.Abort()
//}
//
//func CreateEntityResponse(c *gin.Context, userRaw *interface{}) {
//
//	//var userRaw models.UserRaw
//	err := c.BindJSON(userRaw)
//	if err != nil || userRaw.Id == 0 {
//		if err != nil {
//			log.Println(err)
//		}
//		c.AbortWithStatus(404)
//		return
//	}
//
//	tnx := db.DB.Txn(false)
//	raw, err := tnx.First("user", "id", userRaw.Id)
//	if err != nil || raw != nil {
//		log.Println(err)
//		c.AbortWithStatus(400)
//		return
//	}
//
//	birth_date := time.Unix(int64(userRaw.Birth_date), 0)
//
//	Age, _ := modules.MonthYearDiff(birth_date, time.Now())
//	user := models.User{Id: userRaw.Id, Birth_date: birth_date, Email: userRaw.Email, Gender: userRaw.Gender, Last_name: userRaw.Last_name, Age: Age}
//
//	tnx = db.DB.Txn(true)
//	err = tnx.Insert("user", &user)
//	if err != nil {
//		log.Println(err)
//		c.AbortWithStatus(404)
//		return
//	}
//	tnx.Commit()
//
//	resp := map[string]string{}
//	c.JSON(200, resp)
//
//	c.Abort()
//}
