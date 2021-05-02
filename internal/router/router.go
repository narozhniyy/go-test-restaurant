package router

import (
	"github.com/labstack/echo"
)

func InitRouter() *echo.Echo {
	e := echo.New()

	e.POST("/assign-table", assignTable)
	e.POST("/order/:table", makeOrder)
	e.POST("/bills/:table", getBills)

	return e
}