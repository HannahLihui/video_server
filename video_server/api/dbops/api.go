package dbops

import (
	"database/sql"
	_ "database/sql"
	"github.com/HannahLihui/video_server/video_server/api/defs"
	"github.com/HannahLihui/video_server/video_server/api/utils"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
	_ "time"
)

type VideoInfo struct {
	Id string
	AuthorId int
	Name string
	Displayctime string
}
func AddUserCredential(loginName string, Pwd string) error {
   stmtIns, err:=dbConn.Prepare("insert into users(login_name,pwd) values(?,?)")
   if(err!=nil){
   	return err
   }
   _, err=stmtIns.Exec(loginName, Pwd)
   if(err!=nil){
   	return err
   }
   defer stmtIns.Close()
   return nil
}
func GetUserCredential(loginName string) (string, error) {
    stmtOuts,err:=dbConn.Prepare("select pwd from users where login_name=?")
    if(err !=nil){
    	log.Printf("%s", err)
    	return "",err
	}
    var pwd string
    err=stmtOuts.QueryRow(loginName).Scan(&pwd)
    if err!=nil && err!=sql.ErrNoRows{
    	return "",err
	}
    defer stmtOuts.Close()
    return pwd,nil

}
func DeleteUser(loginName string, pwd string) error  {
	stmtDel, err:= dbConn.Prepare("delete from users where login_name=? and pwd =?");
	if(err!=nil){
		return err
	}
	_,err=stmtDel.Exec(loginName, pwd)
	if(err!=nil){
		return err
	}
	defer stmtDel.Close()
	return nil

}
func AddNewVideo(aid int, name string)(*defs.VideoInfo,error){
    vid, err:=utils.NewUUID()
    if(err!=nil){
    	return nil,err
	}
    t:=time.Now()
    ctime:=t.Format("Jan 02 2006,15:04:05")
    stmIns, err:=dbConn.Prepare("INSERT INTO video_info(id, author_id,name,display_ctime) values (?,?,?,?)")
    if(err!=nil){
    	return nil,err
	}
    _,err=stmIns.Exec(vid,aid,name,ctime)
    if(err!=nil){
    	return nil,err
	}
    res:=&defs.VideoInfo{Id:vid, AuthorId:aid,Name:name, Displayctime:ctime  }
    defer stmIns.Close()
    return res,err
}
func GetVideoInfo(vid string)(*defs.VideoInfo,error){
	stmtOut,err:=dbConn.Prepare("SELECT author_id, name, display_ctime from video_info where vid=?")
	var aid int
	var dct string
	var name string
	err=stmtOut.QueryRow(vid).Scan(&aid,&name,&dct)
	if err!=nil && err!=sql.ErrNoRows {
		return nil,err
	}
	if err == sql.ErrNoRows {
		return nil ,err
	}
	defer stmtOut.Close()
	res:=&defs.VideoInfo{Id:vid, Name:name, AuthorId:aid, Displayctime:dct}
	return res,err
}
func DeleteVideo(vid string) error  {
	stmtDel, err:= dbConn.Prepare("delete from video_info where id=?");
	if(err!=nil){
		return err
	}
	_,err=stmtDel.Exec(vid)
	if(err!=nil){
		return err
	}
	defer stmtDel.Close()
	return nil

}
