package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	model "github.com/wbing441282413/goGormTest/gorm-relation/relation-belongsto/model"
)

var db *gorm.DB

func init() {
	//用户名:密码@tcp(数据库ip或域名:端口)/数据库名称?charset=数据库编码&parseTime=True&loc=Local
	var err error
	db, err = gorm.Open("mysql", "root:12345@tcp(127.0.0.1:3306)/db_01?"+
		"charset=utf8&parseTime=True&loc=Local")
	//有点像go的数据库包一样，使用open方法来两将诶数据库
	if err != nil {
		panic(err)
	}
	// fmt.Printf("%#v\n", db)

	db.LogMode(true) //开启日志打印
	// db.SetLogger(log.New(os.Stdout, "\r\n", 0)) //设置日志格式

	//配置连接池
	db.DB().SetMaxIdleConns(10)  //最大空闲连接池数
	db.DB().SetMaxOpenConns(100) //数据库打开的最大连接数
}

func main() {

	var songs []model.Song
	db.Preload("Singer").Find(&songs) //默认是左连接的
	for _, v := range songs {
		v.ToString()
	}

}
