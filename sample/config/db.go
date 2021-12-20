package config

import (
	"log"
	"os"
)

type Db struct {
	USER   string
	PASS   string
	HOST   string
	PORT   string
	DBNAME string
}

func (c Db) GetDsn() string {
	c = Db{
		USER:   os.Getenv("MYSQL_USER"),
		PASS:   os.Getenv("MYSQL_PASSWORD"),
		HOST:   os.Getenv("DB_HOST"),
		PORT:   os.Getenv("DB_PORT"),
		DBNAME: os.Getenv("MYSQL_DATABASE")}

	host := "tcp(" + c.HOST + ":" + c.PORT + ")"
	dsn := c.USER + ":" + c.PASS + "@" + host + "/" + c.DBNAME + "?charset=utf8mb4&parseTime=true&loc=Asia%2FTokyo"

	log.Printf("DB 接続情報 %s:", dsn)
	return dsn
}
