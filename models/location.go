//go:generate reform

package models

import "encoding/json"

//reform:location
type Location struct {
	Id       uint32 `reform:"id,pk" json:"id"`
	Place    string `reform:"place" json:"place" binding:"required"`
	Country  string `reform:"country" json:"country" binding:"required"`
	City     string `reform:"city" json:"city" binding:"required"`
	Distance uint32 `reform:"distance" json:"distance" binding:"required"`
}

func (u *Location) MarshalJSON() ([]byte, error) {
	type Alias Location
	return json.Marshal(&struct {
		*Alias
	}{
		Alias: (*Alias)(u),
	})
}
