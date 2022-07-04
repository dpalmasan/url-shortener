module github.com/url-shortener

go 1.18

replace github.com/url-shortener/types => ../types

require (
	github.com/sirupsen/logrus v1.8.1
	gorm.io/driver/mysql v1.3.4
	gorm.io/gorm v1.23.7
)

require (
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.4 // indirect
	golang.org/x/sys v0.0.0-20191026070338-33540a1f6037 // indirect
)

replace github.com/url-shortener/db/mysqldb => ../mysqldb
