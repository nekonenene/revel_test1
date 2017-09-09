package models

import (
	"net/url"
	"github.com/revel/revel"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	Gorm *gorm.DB
)

func InitDB() {
	driver := revel.Config.StringDefault("db.driver", "mysql") // MySQLを想定
	username := revel.Config.StringDefault("db.username", "root")
	password := revel.Config.StringDefault("db.password", "")
	protocol := revel.Config.StringDefault("db.protocol", "tcp")
	hostname := revel.Config.StringDefault("db.hostname", "localhost")
	port := revel.Config.StringDefault("db.port", "3306")
	dbName := revel.Config.StringDefault("db.database_name", "revel")
	timezone := revel.Config.StringDefault("db.timezone", "UTC")
	charset := revel.Config.StringDefault("db.charset", "utf8mb4")
	collation := revel.Config.StringDefault("db.collation", "utf8mb4_bin")

	dsn := username + ":" + password +
		"@" + protocol + "(" + hostname + ":" + port + ")" +
		"/" + dbName +
		"?parseTime=True&loc=" + url.PathEscape(timezone) +
		"&charset=" + charset + "&collation=" + collation

	// revel.INFO.Printf("%s", dsn)

	var err error
	Gorm, err = gorm.Open(driver, dsn)
	if err != nil {
		defer Gorm.Close() // panic が起きても必ず閉じる
		panic(err)
	}

	Gorm.LogMode(true)
	Gorm.SetLogger(gorm.Logger{revel.INFO})
	Gorm.SingularTable(true) // テーブル名を単数形に

	autoMigrate()
	revel.INFO.Println("Connected to database.")
}

func autoMigrate() {
	Gorm.Set("gorm:table_options", "ENGINE=InnoDB")
	Gorm.AutoMigrate(&Memo{})
}
