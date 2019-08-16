package dbops

import (
	"strconv"
	"video-server/api/defs"
	"database/sql"
	"sync"
)

func InsertSession(sid string, ttl int64, username string) error {
	ttlstr := strconv.FormatInt(ttl, 10)
	stmtIns, err := dbConn.Prepare("INSERT INTO sessions (session_id, TTL, login_name) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = stmtIns.Exec(sid, ttlstr, username)
	if err != nil {
		return err
	}
	defer stmtIns.Close()
	return nil
}

func RetrieveSession(sid string) (*defs.SimpleSession, error) {
	stmtOut, err := dbConn.Prepare("SELECT TTL, login_name FROM sessions WHERE session_id = ?")
	if err != nil {
		return nil, err
	}
	var ttl string
	var username string
	err = stmtOut.QueryRow(sid).Scan(&ttl, &username)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	ttlint, err := strconv.ParseInt(ttl, 10, 64)
	if err != nil {
		return nil, err
	}
	ss := &defs.SimpleSession{
		Username: username,
		TTL:      ttlint,
	}
	defer stmtOut.Close()
	return ss, nil
}

func RetriveAllSessions() (*sync.Map, error) {
	stmtOut, err := dbConn.Prepare("SELECT * FROM sessions")
	if err != nil {
		return nil, err
	}
	rows, err := stmtOut.Query()
	if err != nil {
		return nil, err
	}
	m := &sync.Map{}
	for rows.Next() {
		var id string
		var ttlstr string
		var username string
		if err := rows.Scan(&id, &ttlstr, &username); err != nil {
			break
		}
		if ttlint, e := strconv.ParseInt(ttlstr, 10, 64); e == nil {
			m.Store(id, &defs.SimpleSession{
				Username: username,
				TTL:      ttlint,
			})
		}
	}
	defer stmtOut.Close()
	return m, nil
}

func DeleteSession(sid string) error {
	stmtDel, err := dbConn.Prepare("DELETE FROM sessions WHERE session_id = ?")
	if err != nil {
		return err
	}
	_, err = stmtDel.Exec(sid)
	if err != nil {
		return err
	}
	return nil
}
