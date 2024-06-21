package sqlites

import "fmt"

func UpFile(filename string) error {
	sqlStr := "INSERT INTO file_info (create_time,filename) VALUES (?,?)"
	createime := TimeNow()
	_, err := db.Exec(sqlStr, createime, filename)
	if err != nil {
		fmt.Println("db inset err :", err)
		return err
	}
	return nil
}
