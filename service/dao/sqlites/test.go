package sqlites

import (
	"fmt"
)

func Dksd() {
	sqlStr := "INSERT INTO text_info (create_time,text) VALUES (?,?)"
	createime := TimeNow()
	_, err := db.Exec(sqlStr, createime, "test")
	if err != nil {
		fmt.Println("db inset err :", err)
	}

}
