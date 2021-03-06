package dbutil

import "fmt"

// Info ...
type Info struct {
	ID   int
	Name string
}

// DBMode ...
// **added for PostgreSQL** - choose "postgres" or "mysql"!
const DBMode = "mysql"

// const DBMode = "postgres"

// DbDriver ...
// const DbDriver = "mysql"
var DbDriver = "mysql" // **added for PostgreSQL**

// User ...
// const User = "root"
var User = "root" // **added for PostgreSQL**

// Password ...
const Password = "tyler"

// DbName ...
const DbName = "go_db1"

// TableName ...
const TableName = "person"

// DataSourceName ...
// dataSourceName := "root:tyler@tcp(127.0.0.1:3306)/go_db1?charset=utf8"
var DataSourceName = fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?charset=utf8",
	User, Password, DbName)
