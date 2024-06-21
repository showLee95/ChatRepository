package sqlites

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var db *sqlx.DB

func Init() (err error) {
	dsn := "./imdata.db"
	db, err = sqlx.Open("sqlite3", dsn)
	if err != nil {
		fmt.Printf("sqlite Init DB failed, err:%v\n", err)
		return
	}
	n, err := tableExists()
	if err != nil {
		fmt.Println("check db teble err:", err)
	}
	if n == false {
		tastCreateTable()
	}
	x, err := tableExists2()
	if err != nil {
		fmt.Println("check db teble err:", err)
	}
	if x == false {
		tastCreateTable2()
	}
	return
}

func Close() {
	db.Close()
}

// check teble
func tableExists() (bool, error) {
	var count int
	query := `SELECT count(*) FROM sqlite_master WHERE type='table' AND name=?`
	err := db.Get(&count, query, "text_info")
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
func tableExists2() (bool, error) {
	var count int
	query := `SELECT count(*) FROM sqlite_master WHERE type='table' AND name=?`
	err := db.Get(&count, query, "file_info")
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

// Create teble
func tastCreateTable() error {
	sqlc := `
	CREATE TABLE "text_info" (
	  "id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	  "text" text(255) NOT NULL,
	  "create_time" integer(4) NOT NULL,
	  "comments" TEXT(255)
	);
	
	CREATE INDEX "indexs"
	ON "text_info" (
	  "text",
	);
    `
	_, err := db.Exec(sqlc)
	if err != nil {
		return err
	}
	return nil
}
func tastCreateTable2() error {
	sqlc := `
	CREATE TABLE "file_info" (
	  "id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	  "filename" text(255) NOT NULL,
	  "create_time" integer(4) NOT NULL,
	  "comments" TEXT(255)
	);
	
	CREATE INDEX "indexs"
	ON "file_info" (
	  "filename",
	);
    `
	_, err := db.Exec(sqlc)
	if err != nil {
		return err
	}
	return nil
}
