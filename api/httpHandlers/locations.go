package httpHandlers

import (
	"github.com/eekrupin/hlc-travels/db"
	"github.com/eekrupin/hlc-travels/models"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

// /locations
func Locations(c *gin.Context) {
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

	filter := ""
	var args []interface{}
	args = append(args, id)
	addConvArgs(&filter, &args, c, "fromDate", RequestIntParam, "and visit.visited_at > ?")
	addConvArgs(&filter, &args, c, "toDate", RequestIntParam, "and visit.visited_at < ?")
	addConvArgs(&filter, &args, c, "gender", RequestStringParam, "and user.gender = ?")
	addConvArgs(&filter, &args, c, "fromAge", RequestAgeParamToTimeUnixViaCurrentTime, "and user.birth_date < ?")
	addConvArgs(&filter, &args, c, "toAge", RequestAgeParamToTimeUnixViaCurrentTime, "and user.birth_date > ?")

	if v, ok := c.Keys["err"]; ok && v == true {
		c.AbortWithStatus(400)
		return
	}

	text := `SELECT 
       cast(avg(visit.mark) as DECIMAL(10,5)) as mark
	from
		visit as visit
	inner join location as location
		on visit.location = location.id
	inner join user as user
		on visit.user = user.id
	where location.id = ?
# 		and visit.visited_at > ?
# 		and visit.visited_at < ?
# 		and user.age > ?
# 		and user.age < ?
# 		and user.gender = ?
		`
	text = text + filter
	var mark float32
	err = db.DB.QueryRow(text, args).Scan(&mark)
	if err != nil {
		log.Println("Error while get mark: ", err.Error())
		c.AbortWithStatus(500)
		return
	}

	//if err != nil {
	//	log.Println(err)
	//	c.AbortWithStatus(404)
	//	return
	//}

	resp := map[string]float32{}
	resp["AVG"] = mark
	c.JSON(200, resp)

	c.Abort()
}

/*func addIntArgs(columns *string, args *[]interface{}, c *gin.Context, request string) {
	if fromDate, ok := RequestIntParam(c, request); ok{
		*columns = *columns + ",fromDate"
		*args = append(*args, fromDate)
	}
}*/
