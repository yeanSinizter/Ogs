package bookingcontroller

import (
	"golang/configs/dbconfig"
	"golang/configs/responseconfig"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetAllBooking(c echo.Context) error {
	booking := []Booking{}

	dbconfig.MySQL.Find(&booking)

	return responseconfig.Handler(c).Success(booking)

}

func CreateDataBooking(c echo.Context) error {
	booking := Booking{}
	err := c.Bind(&booking)
	if err != nil {
		return responseconfig.Handler(c).BadRequest(booking)

	}
	if err = c.Validate(booking); err != nil {
		return responseconfig.Handler(c).BadRequest(booking)

	}
	dbconfig.MySQL.Create(&booking).Debug()
	return responseconfig.Handler(c).Success(booking)

}

func DeleteDataBooking(c echo.Context) error {
	booking := Booking{}
	bookingId := c.Param("bookingId")
	getBookingId, err := strconv.Atoi(bookingId)
	if err != nil {
		return responseconfig.Handler(c).BadRequest(booking)

	}
	booking.Id = getBookingId
	dbconfig.MySQL.Delete(&booking)
	return responseconfig.Handler(c).Success(booking)

}

type Booking struct {
	Id           int    `json:"id"`
	StartDate    string `json:"start_date"`
	EndDate      string `json:"end_date"`
	RoomsId      int    `json:"rooms_id"`
	AmountPerson int    `json:"amount_person" validate:"required,gte=1"`
	UsersId      int    `json:"users_id"`
}
