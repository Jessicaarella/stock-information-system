package main

import (
	routes "alta/delivery/routes/external"
	"alta/pkg/db"

	_ "github.com/go-sql-driver/mysql"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type Stock struct {
	CompanyCode string
	CompanyName string
	LastPrice   int
}

func main() {
	// Initiate echo framework
	e := echo.New()
	log.Info("Initiate echo successfull!")

	// Initiate database
	database, err := db.Init()
	if err != nil {
		log.Errorf("Failed connect to db: %v", err)
		return
	}
	log.Info("Connecting to mysql successfull!")

	// Initiate routes
	route := routes.New(database)
	route.Init(e)
	log.Info("Checking routing successfull!")

	e.Logger.Fatal(e.Start(":1323"))
}
