package sqlites

import (
	"fmt"
)

func Text(text string) {
	sqlStr := "INSERT INTO text_info (create_time,text) VALUES (?,?)"
	createime := TimeNow()
	_, err := db.Exec(sqlStr, createime, text)
	if err != nil {
		fmt.Println("db inset err :", err)
	}

}
