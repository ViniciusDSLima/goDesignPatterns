package main

import "fmt"

type IBDConnection interface {
	connect()
}

type DBConnection struct {
	Db IBDConnection
}

func (conn DBConnection) DbConnect() {
	conn.Db.connect()
}

type MysqlConnection struct {
	ConnectionString string
}

func (conn MysqlConnection) connect() {
	fmt.Println("Mysql " + conn.ConnectionString)
}

type PostgresConnection struct {
	ConnectionString string
}

func (conn PostgresConnection) connect() {
	fmt.Println("Postgres " + conn.ConnectionString)
}

type MongoDbConnection struct {
	ConnectionString string
}

func (conn MongoDbConnection) connect() {
	fmt.Println("Mongo " + conn.ConnectionString)
}

func main() {

	MysqlConnection := MysqlConnection{ConnectionString: "mysql://localhost:27017"}
	con := DBConnection{Db: MysqlConnection}
	con.DbConnect()

	PgConnection := PostgresConnection{ConnectionString: "postgres://localhost:5432/postgres"}
	con2 := DBConnection{Db: PgConnection}
	con2.DbConnect()

	MongoConnection := MongoDbConnection{ConnectionString: "mongodb://localhost:27017"}
	con3 := DBConnection{Db: MongoConnection}
	con3.DbConnect()
}
