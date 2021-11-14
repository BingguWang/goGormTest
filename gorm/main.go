package main

import (
	"fmt"

	// _ "github.com/jinzhu/gorm/dialects/mysql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/wbing441282413/goGorm/gorm/gormutils"
)

//orm的作用主要是映射程序字段和数据库字段，方便持久化
//数据库驱动依赖不能少
func main() {
	//用户名:密码@tcp(数据库ip或域名:端口)/数据库名称?charset=数据库编码&parseTime=True&loc=Local
	db, err := gorm.Open("mysql", "root:12345@tcp(127.0.0.1:3306)/db_01?"+
		"charset=utf8&parseTime=True&loc=Local")
	//有点像go的数据库包一样，使用open方法来两将诶数据库
	if err != nil {
		panic(err)
	}
	// fmt.Printf("%#v\n", db)
	defer db.Close()

	db.LogMode(true) //开启日志打印
	// db.SetLogger(log.New(os.Stdout, "\r\n", 0)) //设置日志格式

	//配置连接池
	db.DB().SetMaxIdleConns(10)  //最大空闲连接池数
	db.DB().SetMaxOpenConns(100) //数据库打开的最大连接数

	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>插入操作>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	s := gormutils.Singer{Name: "周杰伦", NickName: "周董"} //满足驼峰命名方法
	//插入数据
	db.Create(&s)

	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>删除操作>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	//删除数据
	var ss gormutils.Singer
	db.Where("name = ?", "谭咏麟").Delete(&ss)

	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>更新操作>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	//修改数据
	db.Model(&ss).Where("mainid = ?", 8).Update("name", "xxx") //每个Update就是一个update语句
	// UPDATE `singers` SET `name` = 'xxx'  WHERE (mainid = 8)

	//查出来再修改，可以直接用查询语句和save语句
	sg := gormutils.Singer{}
	db.Where("mainid = ?", 3).Take(&sg)
	sg.SetName("tiger1")
	sg.SetNickName("虎子1")
	db.Model(&sg).Where("mainid = ?", sg.MainId).Update(&sg) //where和update方法顺序不能反了

	//updates跟新多个字段
	mp := make(map[string]interface{})
	mp["name"] = "小青"
	mp["nickName"] = "妖怪"
	db.Model(&sg).Where("mainid > 10").Updates(mp)

	//UPDATE singer SET mainid = mainid + 1 WHERE id = '20' //SQL中有表达式的
	db.Model(&sg).Where("mainid = 20").Update("age", gorm.Expr("age + 1"))

	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>事务>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	//类似database/sql包的事务，
	//开启事务
	tx := db.Begin()
	ad := gormutils.Singer{
		MainId:   1, //此主键已存在，会报错，看下是否能回滚
		Name:     "陶喆",
		Age:      51,
		NickName: "DT",
	}
	if err := tx.Create(&ad); err != nil {
		//回滚
		tx.Rollback() //回滚后后面的所有SQL都会失效，包括查询也查不出
		//如果不在这调用rollback，那就是错误的SQL不会执行，但是其他正确的操作会执行，那就没有回滚的效果了，
		// 所以还是要执行rollback才对
		fmt.Println(err)
	}

	var bd gormutils.Singer
	tx.Where("mainid = ?", 2).Take(&bd) //即使是查询操作还是不会操作，因为前面有回滚
	fmt.Println(bd)
	bd.SetName("刘德华")
	tx.Model(&bd).Where("mainid = ?", bd.GetMainId()).Update(&bd)

	//提交事务
	tx.Commit()

}
