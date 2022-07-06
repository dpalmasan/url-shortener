# Url Shortener

A simple implementation of a url shortener service. The service implements two endpoints:

* `POST /shortener`: Receives as input a `url` (in `JSON` body) and creates a hash for the url
* `GET /<hash>`: Based on your host, receives the hash and redirects to the original url that generated the hash

## To Run

The assumption is that you have a SQL instance running. You need to create the database and the tables. For this you can just run (assuming you have `root` access, if not, just use a user that has the required permissions):

`mysql -uroot -p < sql/init_db.sql`

To start the server:

`MYSQL_DSN='user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local' go run cmd/main.go`

By default it uses port `4000`

## TODOs

* Extend to a distributed system
* Avoid duplicated urls when creating a new hash
* Dockerize