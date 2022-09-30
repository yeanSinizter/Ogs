package routers

import (
	"github.com/labstack/echo/v4"
)

type MainRoutesVersion struct {
	ApiV1 echo.Group
	ApiV2 echo.Group
}

func Init(e *echo.Echo) {
	e.Static("public", "../static/public")

	routesVersions := new(MainRoutesVersion)
	routesVersions.ApiV1 = *e.Group("/api/v1")
	routesVersions.ApiV2 = *e.Group("/api/v2")

	routesVersions.regisUserControllerEndPoint()
	routesVersions.regisRoomControllerEndPoint()
	routesVersions.regisBookingControllerEndPoint()
}
