package model

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type Song struct {
	gorm.Model //// 将字段 `ID`, `CreatedAt`, `UpdatedAt`, `DeletedAt` 注入到 `User` 模型中
	//创建的时候会默认把Model中的ID作为主键，想用自己的字段主键加上;"primary_key"就可以了
	SongId   int64  `gorm:"type:int;unique_index;not null;primary_key"`
	SongName string `gorm:"type:varchar(30)"`

	SongList []Song `gorm:"many2many:sheet_songs;"`
}

func (song Song) ToString() {
	fmt.Printf("{songId: %v, songName: %v}\n",
		song.SongId, song.SongName)
}
