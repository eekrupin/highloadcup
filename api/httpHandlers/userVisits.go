package httpHandlers

import (
	"github.com/eekrupin/hlc-travels/db"
	"github.com/eekrupin/hlc-travels/models"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

type UserVisit struct {
	Mark       uint   `json:"mark"`
	Visited_at uint32 `json:"visited_at"`
	Place      string `json:"place"`
}

// /userVisits
func UserVisits(c *gin.Context) {
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

	/*	fromDate - посещения с visited_at > fromDate
		toDate - посещения до visited_at < toDate
		country - название страны, в которой находятся интересующие достопримечательности
		toDistance - возвращать только те места, у которых расстояние от города меньше этого параметра*/

	filter := ""
	var args []interface{}
	args = append(args, id)
	addConvArgs(&filter, &args, c, "fromDate", RequestIntParam, "and visit.visited_at > ?")
	addConvArgs(&filter, &args, c, "toDate", RequestIntParam, "and visit.visited_at < ?")
	addConvArgs(&filter, &args, c, "country", RequestStringParam, "and location.country = ?")
	addConvArgs(&filter, &args, c, "toDistance", RequestIntParam, "and location.distance < ?")

	if v, ok := c.Keys["err"]; ok && v == true {
		c.AbortWithStatus(400)
		return
	}

	text := `SELECT 
       visit.mark as mark,
       visit.visited_at as visited_at,
       visit.place as place
	from
		visit as visit
	inner join location as location
		on visit.location = location.id
	inner join user as user
		on visit.user = user.id
	where location.id = ?
# 		and visit.visited_at > ?
# 		and visit.visited_at < ?
# 		and location.country = ?
# 		and location.distance < ?
		`
	text = text + filter + "/n" + "ORDER BY visited_at"
	rows, err := db.DB.Query(text, args)
	if err != nil {
		log.Println("Error while get UserVisits: ", err.Error())
		c.AbortWithStatus(500)
		return
	}

	var visits []interface{}
	for rows.Next() {
		var userVisit UserVisit
		err = rows.Scan(&userVisit)
		if err != nil {
			log.Println("Error while Scan UserVisit: ", err.Error())
			c.AbortWithStatus(500)
			return
		}
		visits = append(visits, userVisit)
	}

	resp := make(map[string]interface{})
	resp["visits"] = visits
	c.JSON(200, resp)

	c.Abort()
}

//func addConvArgs(filter *string, args *[]interface{}, c *gin.Context, request string, convert func (*gin.Context, string)(interface{}, bool), filterString string) {
//	if fromDate, ok := convert(c, request); ok{
//		*filter = *filter + "/n" + filterString
//		*args = append(*args, fromDate)
//	}
//}

/*func addIntArgs(columns *string, args *[]interface{}, c *gin.Context, request string) {
	if fromDate, ok := RequestIntParam(c, request); ok{
		*columns = *columns + ",fromDate"
		*args = append(*args, fromDate)
	}
}*/
