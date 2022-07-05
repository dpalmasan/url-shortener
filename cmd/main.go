package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
	"github.com/url-shortener/api"
	"github.com/url-shortener/db/mysqldb"
	"github.com/url-shortener/models/providers"
	"github.com/url-shortener/utils"
)

var (
	log = logrus.New()
)

func init() {
	// Logging =================================================================
	// Setup the logger backend using Sirupsen/logrus and configure
	// it to use a custom JSONFormatter. See the logrus docs for how to
	// configure the backend at github.com/Sirupsen/logrus
	log.Formatter = new(logrus.JSONFormatter)

	// Connect to DB
	mysqldb.ConnectToMySql()
}

func main() {
	PORT := utils.Getenv("AUTH_SERVER_PORT", "4000")
	db := providers.MysqlDBUrl{}

	defer func() {
		mysqlDb, err := mysqldb.Session.DB()
		if err != nil {
			panic(err)
		}
		err = mysqlDb.Close()
		if err != nil {
			panic(err)
		}
	}()

	router := chi.NewRouter()

	router.Mount("/shortener", api.Routes(db))
	router.Mount("/", api.RedirectRoutes(db))
	log.Infof("Running service on port %s", PORT)
	http.ListenAndServe(":"+PORT, router)
}
