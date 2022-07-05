package mysqldb

import (
	"github.com/sirupsen/logrus"
	"github.com/url-shortener/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	log          = logrus.New()
	Session      *gorm.DB
	DatabaseName string
)

func init() {
	log.Formatter = new(logrus.JSONFormatter)
}

func ConnectToMySql() {
	DatabaseName = utils.Getenv("MYSQL_DBNAME", "auth")
	dsn := utils.Getenv("MYSQL_DSN", "")

	if dsn == "" {
		log.Panic("MYSQL_DSN env var should be set!")
		panic("MYSQL_DSN is not set!")
	}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic("Failed to connect to MySQL", err)
		panic(err)
	}

	mysqlDb, err := db.DB()
	if err != nil {
		if err = mysqlDb.Ping(); err != nil {
			mysqlDb.Close()
			log.Panic("Failed to ping MySQL DB")
			panic(err)
		}
	}

	log.Info("Connected to MySQL DB!")
	Session = db
}
