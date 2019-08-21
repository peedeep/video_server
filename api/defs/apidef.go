// Data model
package defs

type UserCredential struct {
	Username string `json:"user_name"`
	Pwd      string `json:"pwd"`
}

type SignedUp struct {
	Success bool `json:"success"`
	SessionId string `json:"session_id"`
}

type Login struct {
	Success bool `json:"success"`
	SessionId string `json:"session_id"`
}

type VideoInfo struct {
	Id           string
	AuthorId     int
	Name         string
	DisplayCtime string
}

type Comment struct {
	Id      string
	VideoId string
	Author  string
	Content string
}

type SimpleSession struct {
	Username string //login name
	TTL int64 //检查session是否过期
}