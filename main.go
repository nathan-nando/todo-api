package main

import (
	db "TodoAPI/database"
	"TodoAPI/internal/factory"
	"TodoAPI/internal/http"
	"TodoAPI/pkg/env"
	"flag"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"os"
)

// @title TodoList API
// @version 1.0
// @description API DOCUMENTATION TODOLIST API
// @BasePath /
func main() {
	loadENV()
	db.Init()

	e := echo.New()

	f := factory.NewFactory(GetDBTerminal())
	http.Init(e, f)

	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}

func loadENV() {
	envTerminal := os.Getenv("ENV")
	env.LoadEnv(envTerminal)
	logrus.Info("Choosen ENV : " + envTerminal)
}

func GetDBTerminal() string {
	dbConfiguration := flag.String("dbConfiguration", "TODO_MONGODB", "choose a db")
	flag.Parse()
	return *dbConfiguration
}
