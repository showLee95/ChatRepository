package postgres

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

var pdb *sqlx.DB

func Init() (err error) {
	dsn := fmt.Sprintf("host=%s port=%d  user=%s password=%s dbname=%s sslmode=disable",
		viper.GetString("postgres.host"),
		viper.GetInt("postgres.port"),
		viper.GetString("postgres.user"),
		viper.GetString("postgres.password"),
		viper.GetString("postgres.dbname"),
	)
	// fmt.Println(dsn)
	// fmt.Println(viper.GetInt("postgres.SetMaxOpenConns"))
	pdb, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		fmt.Printf("Postgres Init DB failed, err:%v\n", err)
		return
	}
	pdb.SetMaxOpenConns(viper.GetInt("postgres.SetMaxOpenConns"))
	pdb.SetMaxIdleConns(viper.GetInt("postgres.SetMaxIdleConns"))
	return

}
func Close() {
	if pdb != nil {
		_ = pdb.Close()
	}
}
