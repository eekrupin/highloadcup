//go:generate reform

package models

import (
	"encoding/json"
	"time"
)

//reform:visit
type Visit struct {
	Id         uint32    `reform:"id,pk"`
	Location   uint32    `reform:"location"`
	User       uint32    `reform:"user"`
	Visited_at time.Time `reform:"visited_at"`
	Mark       uint      `reform:"mark"`
}

type VisitRaw struct {
	Id         uint32 `json:"id" binding:"required"`
	Location   uint32 `json:"location" binding:"required"`
	User       uint32 `json:"user" binding:"required"`
	Visited_at uint32 `json:"visited_at" binding:"required"`
	Mark       uint   `json:"mark" binding:"required"`
}

func (u *Visit) MarshalJSON() ([]byte, error) {
	type Alias Visit
	return json.Marshal(&struct {
		Birth_date int64 `json:"visited_at"`
		*Alias
	}{
		Birth_date: u.Visited_at.Unix(),
		Alias:      (*Alias)(u),
	})
}

func VisitFromRaw(visitRaw *VisitRaw) (user *Visit) {
	visited_at := time.Unix(int64(visitRaw.Visited_at), 0)

	visit := &Visit{Id: visitRaw.Id, Location: visitRaw.Location, User: visitRaw.User, Visited_at: visited_at, Mark: visitRaw.Mark}
	return visit
}
