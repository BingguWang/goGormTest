package gormutils

import (
	"time"
)

type Song struct {
	// gorm.Model //// 将字段 `ID`, `CreatedAt`, `UpdatedAt`, `DeletedAt` 注入到 `User` 模型中
	//创建的时候会默认把Model中的ID作为主键，想用自己的字段主键加上;"primary_key"就可以了
	SongId      int64  `gorm:"type:int;unique_index;not null;primary_key"`
	SongName    string `gorm:"type:varchar(30)"`
	SingerId    uint   `gorm:"type:int unsigned"`
	ListenCount uint   `gorm:"type:int unsigned;default:0"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time `sql:"index"`
}

// 只要实现Tabler接口就能自定义表名
func (*Song) TableName() string {
	return "song" // 返回你要自定义的表名
}
