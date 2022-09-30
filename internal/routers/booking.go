package routers

import "golang/internal/controllers/bookingcontroller"

func (mainRoutesVersion MainRoutesVersion) regisBookingControllerEndPoint() {
	booking := mainRoutesVersion.ApiV1.Group("/booking")
	booking.GET("", bookingcontroller.GetAllBooking)
	booking.POST("", bookingcontroller.CreateDataBooking)
	booking.DELETE("/:bookingId", bookingcontroller.DeleteDataBooking)
}
