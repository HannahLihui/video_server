package dbops

import (
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
}
func testAddUser(t *testing.T)  {
	err:=AddUserCredential("lihh", "2333")
	if(err!=nil){
		t.Errorf("error of adduser")
	}
}
func testGetUser(t *testing.T)  {
	pwd, err:=GetUserCredential("lihh")
	if(pwd!="2333"|| err!=nil){
		t.Errorf("get user error")
	}

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