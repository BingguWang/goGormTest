package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	model "github.com/wbing441282413/goGorm/gorm-select/model"
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
	s := model.Song{}

	/*
		为什么没有指定表名，怎么映射结构体和表的呢》会自动把结构体名+s作为对应的表名
	*/

	// db.Find(&s) //获取所有没删除的记录
	// db.First(&s) //按住键递增排序，获第一条数据
	// db.Take(&s) //不指定排序，获取一条数据
	// db.Last(&s) //按主键递减排序，获取第一条数据
	// db.First(&s, 10) //按主键查询，仅限主键是数字的表,主键递增提一条数据

	//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>where

	// db.Where("singer_id = ?", 1).First(&s)
	// db.Where("singer_id = ?", 1).Find(&s)
	// db.Where("song_id in (?)", []string{"1", "2"}).Find(&s)
	// db.Where("song_name like ?", "%ww%").Find(&s)
	// db.Where("song_name like (?) AND listen_count >= ?", "%ww%", "0").Find(&s)

	// tody := time.Now() //都可以
	// // tody := time.Now().Format("2006-01-02 15:04:05")
	// db.Where("created_at <= ?", tody).Find(&s)

	// d, _ := time.ParseDuration("-24h")
	// lastday := tody.Add(d)
	// db.Where("created_at between ? and ?", lastday, tody).Find(&s)

	// //where可以传结构体
	// db.Where(&model.Song{SongName: "ww", SongId: 1}).Find(&s) //当含有传入的结构体体含零值的时候，含零值的条件会被忽略

	// //where可以传map
	// db.Where(map[string]interface{}{"song_name": "ww", "song_id": 2}).Find(&s)

	//Not，和 Where查询类似
	// db.Not("song_id", 1).Find(&s)                //条件相当于`song_id` NOT IN (1)
	// db.Not("song_id", []int{1, 2}).Find(&s)      //条件相当于`song_id` NOT IN (1,2)
	// db.Not(&model.Song{SongName: "ww"}).Find(&s) //song_name` <> 'ww'

	// db.Where("song_name = ? or song_id = 1", "ww").Find(&s) //((song_name = 'ww' or song_id = 1))
	// db.Where("song_name = ?", "ww").Or("song_id = 1").Find(&s) //((song_name = 'ww') OR (song_id = 1))
	// db.Where(&model.Song{SongName: "ww"}).Or(&model.Song{SongId: 1}).Find(&s)

	// db.Find(&s, &model.Song{SongName: "ww", SingerId: 1})
	// db.Find(&s, map[string]interface{}{"Song_name": "ww", "Song_id": 1})

	// db.Find(&s, 1) //where 主键=1

	// 为查询 SQL 添加额外的选项
	// db.Set("gorm:query_option", "FOR UPDATE").First(&s, 1) //select ....for update

	//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>排序order
	// db.Order("singer_id desc").Order("song_name").Find(&s) //ORDER BY singer_id desc,song_name
	// db.Order("singer_id desc").Find(&s).Order("song_name").Find(&s)       //链式查询，相当于两次查询，第一次查询是第二次查询的范围基础
	// db.Order("singer_id desc").Find(&s).Order("song_name", true).Find(&s) /*和上一句不同是第二次查询的排序不同
	// 上面是ORDER BY singer_id desc,song，下面加了true的是 ORDER BY song_name,
	// */

	//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>FirstOrInit,        仅适用与于 struct 和 map 条件

	// db.FirstOrInit(&s, model.Song{SongName: "不存在"}) //会先按`song_name` = '不存在'找，找到就找到，找不到就按写的内容赋值给s
	// db.FirstOrInit(&s, map[string]interface{}{"song_name": "xxxx"})
	// db.Where(model.Song{SongName: "ddd"}).FirstOrInit(&s)

	//Attrs，查不到就会按照Attrs赋值给结构体
	// db.Where(model.Song{SongName: "ww"}).Attrs(model.Song{SongName: "xxx", ListenCount: 30}).FirstOrInit(&s)

	//Assign
	//与FirstOrInit不同，不查不查得到都会按Assign中的属性赋值给结构体
	// db.Where(model.Song{SongId: 1}).Assign(model.Song{SongName: "ww"}).FirstOrInit(&s)

	//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>FirstOrCreate	  仅适用与于 struct 和 map 条件
	//与FirstOrInit类似，但是FirstOrCreate会插入数据库，找到了就更新到数据库，找不到就查到数据库，其他和FirstOrInit一样

	//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>子查询
	// db.Where("listen_count > (?)", db.Table("songs").Select("AVG(listen_count)").Where("singer_id = ?", "0").QueryExpr()).Find(&s)

	//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>Limit
	// 用 -1 取消 LIMIT 限制条件
	// db.Limit(10).Find(&s).Limit(-1).Find(&s) //链式，相当于查了两次，不加-1，第二次也会有limit限制

	// var ss []model.Song
	//>>>>>>>>>>>>>>>>>>>>>>>>>>>>	查询指定字段 	select
	// db.Select("song_name, song_id").Where("song_name like (?)", "%ww%").Find(&ss)

	//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>  Count count要放在语句最后,直接count，不需要find,但是不用find，就需要指定表
	// var count int
	// db.Table("songs").Where("song_name like (?)", "%ww%").Count(&count)
	// fmt.Println(count)

	//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>		group by  having  row   scan
	// rows, _ := db.Table("songs").Select("singer_id , sum(song_id) as total").Group("singer_id").Having("singer_id = 1").Rows()
	// var total int
	// defer rows.Close()
	// for rows.Next() {
	// 	rows.Scan(&s.SingerId, &total) //rows和scan一起用
	// 	fmt.Printf("singerid:%v	total:%v\n", s.SingerId, total)
	// }

	//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>	连接查询
	// rows, _ := db.Table("songs").Select("song_id,song_name,si.name as singer_name").Joins("left join singers si on si.mainid = songs.singer_id").Rows()
	// var songid int
	// var songname, singername string
	// for rows.Next() {
	// 	rows.Scan(&songid, &songname, &singername)
	// 	fmt.Printf("%v,%v,%v\n", songid, songname, singername)
	// }

	/*
		注意所有使用链式查询的地方，一个链中前面的查询是后面查询的基础

	*/
	fmt.Printf("%#v\n", s)
}
