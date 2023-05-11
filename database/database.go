package database

import (
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"strings"
)

var dbConnections map[string]interface{}

func Init() {
	dbConfigurations := map[string]Db{
		"TODO_MYSQL": &dbMySQL{
			db: db{
				Host: os.Getenv("DB_HOST_MYSQL"),
				User: os.Getenv("DB_USER_MYSQL"),
				Pass: os.Getenv("DB_PASS_MYSQL"),
				Port: os.Getenv("DB_PORT_MYSQL"),
				Name: os.Getenv("DB_NAME_MYSQL"),
			},
			Charset:   "utf8mb4",
			ParseTime: "True",
			Loc:       "Local",
		},
		//"TODO_POSTGRES": &dbPostgres{
		//	db: db{
		//		Host: os.Getenv("DB_HOST_POSTGRES"),
		//		User: os.Getenv("DB_USER_POSTGRES"),
		//		Pass: os.Getenv("DB_PASS_POSTGRES"),
		//		Port: os.Getenv("DB_PORT_POSTGRES"),
		//		Name: os.Getenv("DB_NAME_POSTGRES"),
		//	},
		//	SslMode: "",
		//	Tz:      "",
		//},
		//"TODO_MONGODB": &dbMongoDB{
		//	db{
		//		Host: os.Getenv("DB_HOST_MONGODB"),
		//		User: os.Getenv("DB_USER_MONGODB"),
		//		Pass: os.Getenv("DB_PASS_MONGODB"),
		//		Port: os.Getenv("DB_PORT_MONGODB"),
		//		Name: os.Getenv("DB_NAME_MONGODB"),
		//	},
		//},
	}

	dbConnections = make(map[string]interface{})

	for k, v := range dbConfigurations {
		db, err := v.Init()
		if err != nil {
			panic(fmt.Sprintf("Failed to connect to database %s", k))
		}
		dbConnections[k] = db
		logrus.Info(fmt.Sprintf("Successfully connected to database %s", k))
	}
}

func Connection(name string) (interface{}, error) {
	if dbConnections[strings.ToUpper(name)] == nil {
		return nil, errors.New("connections is undefined")
	}
	return dbConnections[strings.ToUpper(name)], nil
}
