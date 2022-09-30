package routers

import "golang/internal/controllers/roomcontroller"

func (mainRoutesVersion MainRoutesVersion) regisRoomControllerEndPoint() {
	room := mainRoutesVersion.ApiV1.Group("/room")
	room.GET("", roomcontroller.GetAllRoom)
	room.GET("/:roomId", roomcontroller.GetRoomParam)
	room.POST("", roomcontroller.CreateDataRoom)
	room.PUT("/:roomId", roomcontroller.UpdateDataRoom)
	room.DELETE("/:roomId", roomcontroller.DeleteDataRoom)

}
