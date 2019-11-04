//go:generate reform

package models

import "time"

//reform:visit
type Visit struct {
	Id         uint32    `reform:"id,pk"`
	Location   uint32    `reform:"location"`
	User       uint32    `reform:"user"`
	Visited_at time.Time `reform:"visited_at"`
	Mark       uint      `reform:"mark"`
}
