package dbops

import (
	"database/sql"
	_ "database/sql"
	_ "github.com/HannahLihui/video_server/video_server/api/defs"
	_ "github.com/HannahLihui/video_server/video_server/api/utils"
	"log"
	"strconv"
	"sync"
	_ "sync"
)

func InsertSession(sid string, ttl int64, uname string) error {
	ttlstr := strconv.FormatInt(ttl, 10)
	stmtIns, err := dbConn.Prepare("INSERT INTO sessions (session_id, ttl, login_name) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = stmtIns.Exec(sid, ttlstr, uname)
	if err != nil {
		return nil
	}
	defer stmtIns.Close()
	return nil
}
func RetriveSession(sid string) (*defs.SimpleSession, error) {
	ss := &defs.SimpleSession{}
	stmtOut, err := dbConn.Prepare("SELECT ttl, login_name FROM sessions WHERE session_id = ?")
	if err != nil {
		return nil, err
	}
	var ttl string
	var uname string
	stmtOut.QueryRow(sid).Scan(&ttl, &uname)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	if res, err := strconv.ParseInt(ttl, 10, 64); err == nil {
		ss.TTL = res
		ss.UserName = uname
	} else {
		return nil, err
	}
	defer stmtOut.Close()
	return ss, nil
}
func RetrieveAllSessions() (*sync.Map, error) {
	m := &sync.Map{}
	stmtOut, err := dbConn.Prepare("SELECT * FROM sessions")
	if err != nil {
		log.Printf("%s", err)
		return nil, err
	}
	rows, err := stmtOut.Query()
	if err != nil {
		log.Panicf("%s", err)
		return nil, err
	}
	for rows.Next() {
		var id string
		var ttlstr string
		var login_time string
		if err := rows.Scan(&id, &ttlstr, &login_time); err != nil {
			log.Panicf("retriebe sessions error:%s", err)
			break
		}
		if ttl, err := strconv.ParseInt(ttlstr, 10, 64); err != nil {
			ss := &defs.SimpleSession{UserName: login_time, TTL: ttl}
			m.Store(id, ss)
		}
	}
	return m, nil
}
func DeletedSession(sid string) error {
	stmtOut, err := dbConn.Prepare("DELETE FROM sessions WHERE session_id=?")
	if err != nil {
		log.Panicf("%s", err)
		return err
	}
	if _, err := stmtOut.Query(sid); err != nil {
		return err
	}
	return nil
}
