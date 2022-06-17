package pojo

import (
	"fmt"
	"time"
)

type Song struct {
	SongId      int       `json:"songId" sql:"song_id"`
	SongName    string    `json:"songName" sql:"song_name"`
	SingerId    int       `json:"singerId" sql:"singer_id"`
	ListenCount int       `json:"listenCount" sql:"listen_count"`
	CreateAt    time.Time `json:"createAt" sql:"created_at"`
	UpdateAt    time.Time `json:"updateAt" sql:"updated_at"`
	DeleteAt    time.Time
}

func (s Song) ToString() {
	fmt.Println(s.SongId, s.SongName, s.SingerId, s.ListenCount, s.CreateAt, s.UpdateAt, s.DeleteAt)
}
