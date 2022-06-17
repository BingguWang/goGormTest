package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/wbing441282413/goGormTest/databasesqlTest/pojo"
	"log"
)

func main() {
	db, err := sql.Open("mysql", "root:12345@tcp(127.0.0.1:3306)/db_01")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("select song_id, song_name, listen_count, created_at from songs")
	defer rows.Close()

	fmt.Printf("%s \n", rows)

	song := pojo.Song{}
	for rows.Next() {
		rows.Scan(&song.SongId, &song.SongName, &song.ListenCount, &song.CreateAt)
		song.ToString()
	}
}
