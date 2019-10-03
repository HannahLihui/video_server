package dbops

import (
	_ "fmt"
	_ "github.com/HannahLihui/video_server/video_server/api/defs"
	_ "github.com/HannahLihui/video_server/video_server/api/utils"
	"testing"
	_ "testing"
)
func clearTables(){
	dbConn.Exec("truncate Users")
}
func TestMain(m *testing.M)  {
	clearTables()
	m.Run()
	clearTables()
}
func TestUserWorkFlow(t *testing.T){
	t.Run("add", testAddUser)
	t.Run("get", testGetUser)
	t.Run("delete", testDeleteUser)
	t.Run("get", testRegetUser)
	t.Run("addV", testAddVideo)

}
func testAddUser(t *testing.T)  {
	err:=AddUserCredential("lihh", "2333")
	if(err!=nil){
		t.Errorf("error of adduser")
	}
}
func testAddVideo(t *testing.T)  {
	_,err:=AddNewVideo(2333, "2333")
	if(err!=nil){
		t.Errorf("error of add video")
	}
}
func testGetVideo(t *testing.T)  {
	pwd, err:=GetUserCredential("lihh")
	if(pwd!="2333"|| err!=nil){
		t.Errorf("get user error")
	}

}
func testGetUser(t *testing.T)  {


}
func testDeleteUser(t *testing.T)  {
	err:=DeleteUser("lihh", "2333")
	if(err!=nil){
		t.Errorf("delete user error")
	}
}
func testRegetUser(t *testing.T)  {
	pwd, err:=GetUserCredential("lihh")
	if(pwd =="2333"){
		t.Errorf("get user error %v",err)
	}
}