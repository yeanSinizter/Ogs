package routers

import (
	"golang/internal/controllers/usercontroller"
)

func (mainRoutesVersion MainRoutesVersion) regisUserControllerEndPoint() {
	user := mainRoutesVersion.ApiV1.Group("/user")
	user.GET("", usercontroller.GetUser)
	user.GET("/:userId", usercontroller.GetUserParam)
	user.POST("", usercontroller.CreateUser)
	user.PUT("/:userId", usercontroller.UpdateData)
	user.DELETE("/:userId", usercontroller.GetUser)
}
