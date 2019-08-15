// 数据库操作
package dbops

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
	"video-server/api/defs"
	"video-server/api/utils"
)

func AddUserCredential(loginName string, pwd string) error {
	stmtIns, err := dbConn.Prepare("INSERT INTO users (login_name, pwd) VALUES (?, ?)")
	if err != nil {
		log.Printf("Add user error: %s", err)
		return err
	}
	_, err = stmtIns.Exec(loginName, pwd)
	if err != nil {
		return err
	}
	defer stmtIns.Close()
	return nil
}

func GetUserCredentail(loginName string) (string, error) {
	stmtOut, err := dbConn.Prepare("SELECT pwd FROM users WHERE login_name = ?")
	if err != nil {
		log.Printf("Get user error: %s", err)
		return "", err
	}
	var pwd string
	err = stmtOut.QueryRow(loginName).Scan(&pwd)
	if err != nil && err != sql.ErrNoRows {
		return "", err
	}
	defer stmtOut.Close()
	return pwd, nil
}

func DeleteUser(loginName string, pwd string) error {
	stmtDel, err := dbConn.Prepare("DELETE FROM users WHERE login_name=? AND pwd = ?")
	if err != nil {
		log.Printf("Delete user error: %s", err)
		return err
	}
	_, err = stmtDel.Exec(loginName, pwd)
	if err != nil {
		return err
	}
	defer stmtDel.Close()
	return nil
}

func AddNewVideo(authorId int, videoName string) (*defs.VideoInfo, error) {
	vid, err := utils.NewUUID()
	if err != nil {
		return nil, err
	}
	t := time.Now()
	ctime := t.Format("Jan 02 2006, 15:04:05") // M D y, HH:MM:SS
	stmtIns, err := dbConn.Prepare("INSERT INTO videos (id, author_id, name, display_ctime) VALUES (?, ?, ?, ?)")
	if err != nil {
		return nil, err
	}
	_, err = stmtIns.Exec(vid, authorId, videoName, ctime)
	if err != nil {
		return nil, err
	}
	videoInfo := &defs.VideoInfo{
		Id:           vid,
		AuthorId:     authorId,
		Name:         videoName,
		DisplayCtime: ctime,
	}
	defer stmtIns.Close()
	return videoInfo, nil
}

func GetVideoInfo(vid string) (*defs.VideoInfo, error) {
	stmtOut, err := dbConn.Prepare("SELECT author_id, name, display_ctime FROM videos WHERE id = ?")
	if err != nil {
		return nil, err
	}
	var aid int
	var name string
	var dct string
	err = stmtOut.QueryRow(vid).Scan(&aid, &name, &dct)
	if err != nil && err != sql.ErrNoRows {
		return nil, nil
	}
	defer stmtOut.Close()
	videoInfo := &defs.VideoInfo{
		Id:           vid,
		AuthorId:     aid,
		Name:         name,
		DisplayCtime: dct,
	}
	return videoInfo, nil
}

func DeleteVideoInfo(vid string) error {
	stmtDel, err := dbConn.Prepare("DELETE FROM videos WHERE id = ?")
	if err != nil {
		return err
	}
	_, err = stmtDel.Exec(vid)
	if err != nil {
		return err
	}
	defer stmtDel.Close()
	return nil
}

func AddNewComments(vid string, aid int, content string) error {
	cid, err := utils.NewUUID()
	if err != nil {
		return err
	}
	stmtIns, err := dbConn.Prepare("INSERT INTO comments (id, video_id, author_id, content) values (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = stmtIns.Exec(cid, vid, aid, content)
	if err != nil {
		return err
	}
	defer stmtIns.Close()
	return nil
}

func ListComments(vid string, from, to int) ([]*defs.Comment, error) {
	stmtOut, err := dbConn.Prepare(
		`SELECT comments.id, users.login_name, comments.content 
				FROM comments
			  	INNER JOIN users ON comments.author_id = users.id
			  	WHERE comments.video_id = ? AND comments.time > FROM_UNIXTIME(?) AND comments.time <= FROM_UNIXTIME(?)`)
	if err != nil {
		return nil, err
	}
	var comments []*defs.Comment
	rows, err := stmtOut.Query(vid, from, to)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var id, name, content string
		if err := rows.Scan(&id, &name, &content); err != nil {
			return comments, err
		} else {
			comments = append(comments, &defs.Comment{
				Id:      id,
				VideoId: vid,
				Author:  name,
				Content: content,
			})
		}
	}
	defer stmtOut.Close()
	return comments, nil
}
