package model

import (
	"github.com/jinzhu/gorm"
)

//歌单，和song是多对多的关系
type Sheet struct {
	gorm.Model //// 将字段 `ID`, `CreatedAt`, `UpdatedAt`, `DeletedAt` 注入
	//创建的时候会默认把Model中的ID作为主键，想用自己的字段主键加上;"primary_key"就可以了

	SheetName string `gorm:"type:varchar(30)"`
	SongList  []Song `gorm:"many2many:sheet_songs;references:songId"` //使用 AutoMigrate 为 User 创建表时，GORM 会自动创建连接表sheet_songs
}
