//go:generate reform

package models

import (
	"encoding/json"
	"time"
)

//reform:user
type User struct {
	Id         uint32    `reform:"id,pk" json:"id"`
	Email      string    `reform:"email" json:"email"`
	First_name string    `reform:"first_name" json:"first_name"`
	Last_name  string    `reform:"last_name" json:"last_name"`
	Gender     string    `reform:"gender" json:"gender"`
	Birth_date time.Time `reform:"birth_date" json:"birth_date"`
	Age        int       //`reform:"age" json:"-"`
}

type UserRaw struct {
	Id         uint32 `json:"id"`
	Email      string `json:"email" binding:"required,emailCheck"`
	First_name string `json:"first_name" binding:"required"`
	Last_name  string `json:"last_name" binding:"required"`
	Gender     string `json:"gender" binding:"required"`
	Birth_date int32  `json:"birth_date" binding:"required"`
}

func (u *User) MarshalJSON() ([]byte, error) {
	type Alias User
	return json.Marshal(&struct {
		Birth_date int64 `json:"birth_date"`
		*Alias
	}{
		Birth_date: u.Birth_date.Unix(),
		Alias:      (*Alias)(u),
	})
}

func UserFromRaw(userRaw *UserRaw) (user *User) {
	birth_date := time.Unix(int64(userRaw.Birth_date), 0)
	//Age, _ := modules.MonthYearDiff(birth_date, time.Now())
	user = &User{Id: userRaw.Id, Birth_date: birth_date, Email: userRaw.Email, Gender: userRaw.Gender, Last_name: userRaw.Last_name, First_name: userRaw.First_name}
	return user
}
