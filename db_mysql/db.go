package db_mysql

import (
	"HelloBeego/models"
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"fmt"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

/*
链接数据库

*/

func Connect(){

	//项目配置
	config := beego.AppConfig
	dbDriver := config.String("db_driverName")
	dbUser := config.String("db_user")         //数据可用户名
	dbPassword := config.String("da_password") //密码
	dbIP := config.String("db_ip")
	dbName := config.String("db_name")
	fmt.Println(dbDriver,dbUser,dbPassword)
	

	//链接数据库
	connUrl := dbUser +":" + dbPassword + "@tcp("+dbIP+")/"+dbName+"?charset=utf8"
	db, err := sql.Open(dbDriver,connUrl)
	if err != nil {//err不为nil,表示连接数据库时出现了错误，程序就在此中断，不用再执行了
		//早发现，早解决
		panic("数据可链接错误，请检查配置")
	}
	Db = db
	fmt.Println(db)
}
//将用户信息保存到数据库中去得函数

func AddUser(u models.User)(int64, error)  {
	//1·将密码进行哈希计算，得到密码hash值，然后再存
	md5Hash := md5.New()
	md5Hash.Write([]byte(u.Password))
	psswordBytes :=md5Hash.Sum(nil)
	u.Password = hex.EncodeToString(psswordBytes)
	//execute，.exe
	result, err := Db.Exec("insert into user(name,birthday,address,password)"+
		"values(?,?,?,?)",u.Name,u.Birthday,u.Address, u.Password)
	if err != nil {
		return -1, err
	}
	row,err :=result.RowsAffected()
	if err != nil {
		return-1, err
	}
	return row,nil
}