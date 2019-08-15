package dbops

import (
	"database/sql"
)

var (
	dbConn *sql.DB
	err    error
)

func init() {
	dbConn, err = sql.Open("mysql",
		"root:Xmima624!@tcp(127.0.0.1:3306)/video_server")
	if err != nil {
		panic(err.Error())
	}
}
