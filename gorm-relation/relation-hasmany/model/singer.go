package model

import "fmt"

type Singer struct {
	MainId   int    `json:"id" gorm:"column:mainid;primary_key"` //``中的是解析JSON时候用的名字说明
	Name     string `json:"name"`
	NickName string `json:"nickName"`
	Age      int    `json:"age"`

	//has many 一对多,foreignKey:对方表中的关联关键字;
	SongList []Song `json:"songList" gorm:"foreignKey:SingerId;"`
}

func (singer *Singer) GetName() string {
	return singer.Name
}

func (singer *Singer) GetMainId() int {
	return singer.MainId
}
func (singer *Singer) GetNickName() string {
	return singer.NickName
}
func (singer *Singer) SetName(name string) {
	singer.Name = name
}

func (singer *Singer) SetMainId(mainid int) {
	singer.MainId = mainid
}

func (singer *Singer) SetNickName(nickName string) {
	singer.NickName = nickName
}
func (s Singer) ToString() {
	fmt.Printf("{singerId: %v, songName: %v, nickname:%v,age:%v}\n",
		s.MainId, s.Name, s.NickName, s.Age)
}
