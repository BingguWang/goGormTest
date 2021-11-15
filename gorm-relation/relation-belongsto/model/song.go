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
	SingerId uint   `gorm:"type:int unsigned"`

	//假设歌曲和歌手是belongs to 的属于关系, song belongs to singer
	// foreignkey:自己表中的关联字段;references:对方表中的关联字段
	Singer Singer `json:"singer" gorm:"foreignkey:SingerId;references:MainId"`
}

func (song Song) ToString() {
	fmt.Printf("{songId: %v, songName: %v, singerId:%v, singer:{singerName:%s, nickName:%s}}\n",
		song.SongId, song.SongName, song.SingerId, song.Singer.GetName(), song.Singer.GetNickName())
}
