package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
)

type Db interface {
	Init() (interface{}, error)
}

type db struct {
	Host string
	User string
	Pass string
	Port string
	Name string
}

type dbMongoDB struct {
	db
}

type dbMySQL struct {
	db
	Charset   string
	ParseTime string
	Loc       string
}

type dbPostgres struct {
	db
	SslMode string
	Tz      string
}

func (d *dbMongoDB) Init() (interface{}, error) {
	//TODO implement me
	panic("implement me")
}

func (d *dbPostgres) Init() (interface{}, error) {
	//TODO implement me
	panic("implement me")
}

func (d *dbMySQL) Init() (interface{}, error) {
	dbInfo := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%s&loc=%s", d.User, d.Pass, d.Host, d.Port, d.Name, d.Charset, d.ParseTime, d.Loc)
	mysqldb, err := sql.Open("mysql", dbInfo)
	if err != nil {
		return nil, err
	}
	db := bun.NewDB(mysqldb, mysqldialect.New())
	return db, nil
}
