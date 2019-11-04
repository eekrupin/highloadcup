//go:generate reform

package models

import "time"

//reform:person
type Person struct {
	Id         uint32    `reform:"id,pk"`
	Email      string    `reform:"email"`
	First_name string    `reform:"first_name"`
	Last_name  string    `reform:"last_name"`
	Gender     string    `reform:"gender"`
	Birth_date time.Time `reform:"birth_date"`
}
