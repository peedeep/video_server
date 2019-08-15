// Data model
package defs

type UserCredential struct {
	Username string `json:"user_name"`
	Pwd      string `json:"pwd"`
}

type VideoInfo struct {
	Id           string
	AuthorId     int
	Name         string
	DisplayCtime string
}

type Comments struct {
	Id       string
	VideoId  string
	AuthorId int
	Content  string
}
