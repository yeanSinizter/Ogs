package configs

import (
	"golang/configs/dbconfig"
	envconfig "golang/configs/envConfig"
	"golang/configs/responseconfig"
	"golang/configs/validatorconfig"
	"golang/internal/routers"

	"github.com/labstack/echo/v4"
)

func Init(e *echo.Echo) {
	envconfig.Init()
	responseconfig.Init(e)
	dbconfig.Init()
	validatorconfig.Init(e)
	routers.Init(e)
}
