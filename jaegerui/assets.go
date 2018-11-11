package jaegerui

import (
	"github.com/labstack/echo"
	"github.com/rai-project/web/jaegerui/static"
)

func Handle() echo.HandlerFunc {
	return echo.WrapHandler(static.Handler)
}
