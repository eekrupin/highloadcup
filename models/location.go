//go:generate reform

package models

//reform:location
type Location struct {
	Id       uint32 `reform:"id,pk"`
	Place    string `reform:"place"`
	Country  string `reform:"country"`
	City     string `reform:"city"`
	Distance uint32 `reform:"distance"`
}
