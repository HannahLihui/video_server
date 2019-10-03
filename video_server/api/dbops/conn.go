package dbops
import (
	"database/sql"
	_ "database/sql"
	_ "github.com/go-sql-driver/mysql"
)
var(
	dbConn * sql.DB
	err error
)
func init(){
	dbConn, err=sql.Open("mysql", "root:lihonghui2333@/video");
	if(err!=nil){
		panic(err.Error())
	}
}
