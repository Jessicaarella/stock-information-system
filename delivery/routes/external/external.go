package routes

import (
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

type Rest struct {
	DB *sqlx.DB
}

func New(DB *sqlx.DB) *Rest {
	return &Rest{
		DB: DB,
	}
}

func (rest *Rest) Init(e *echo.Echo) {
	e.GET("/users", rest.controllerGetAllUsers)
	e.GET("/user/:id", rest.controllerGetUserById)
	e.POST("/register", rest.controllerRegister)
}
