package roomcontroller

import (
	"golang/configs/dbconfig"
	"golang/configs/responseconfig"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetAllRoom(c echo.Context) error {
	rooms := []Rooms{}

	dbconfig.MySQL.Find(&rooms)

	return responseconfig.Handler(c).Success(rooms)
}

func GetRoomParam(c echo.Context) error {
	rooms := Rooms{}
	roomsId := c.Param("roomId")
	getUserId, err := strconv.Atoi(roomsId)
	if err != nil {
		return responseconfig.Handler(c).BadRequest(rooms)
	}
	rooms.Id = getUserId
	dbconfig.MySQL.Find(&rooms)
	return responseconfig.Handler(c).Success(rooms)
}

func CreateDataRoom(c echo.Context) error {
	room := Rooms{}
	err := c.Bind(&room)
	if err != nil {
		return responseconfig.Handler(c).BadRequest(room)

	}
	dbconfig.MySQL.Create(&room).Debug()
	return responseconfig.Handler(c).Success(room)
}

func UpdateDataRoom(c echo.Context) error {
	room := Rooms{}
	roomId := c.Param("roomId")
	err := c.Bind(&room)
	if err != nil {
		return responseconfig.Handler(c).BadRequest(room)

	}
	getRoomId, err := strconv.Atoi(roomId)
	if err != nil {
		return c.String(400, err.Error())
	}
	room.Id = getRoomId
	dbconfig.MySQL.Updates(&room).Debug()
	return responseconfig.Handler(c).Success(room)

}

func DeleteDataRoom(c echo.Context) error {
	room := Rooms{}
	roomId := c.Param("roomId")
	getRoomId, err := strconv.Atoi(roomId)
	if err != nil {
		return responseconfig.Handler(c).BadRequest(room)

	}
	room.Id = getRoomId
	dbconfig.MySQL.Delete(&room).Debug()
	return responseconfig.Handler(c).Success(room)

}

type Rooms struct {
	Id            int    `json:"id"`
	RoomName      string `json:"room_name"`
	MaximumPerson int    `json:"maximum_person"`
}
