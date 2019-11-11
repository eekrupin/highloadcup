//go:generate reform

package models

import "time"

//reform:user
type User struct {
	Id         uint32    `reform:"id,pk"`
	Email      string    `reform:"email"`
	First_name string    `reform:"first_name"`
	Last_name  string    `reform:"last_name"`
	Gender     string    `reform:"gender"`
	Birth_date time.Time `reform:"birth_date"`
	Age        int       `reform:"age"`
}

type UserRaw struct {
	Id         uint32 `json:"id"`
	Email      string `json:"email" binding:"required,emailCheck"`
	First_name string `json:"first_name" binding:"required"`
	Last_name  string `json:"last_name" binding:"required"`
	Gender     string `json:"gender" binding:"required"`
	Birth_date uint32 `json:"birth_date" binding:"required"`
}
