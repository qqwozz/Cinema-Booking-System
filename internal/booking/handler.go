package booking

import (
	"net/http"

	"github.com/qqwozz/Cinema-Booking-System/internal/utils"
)

type handler struct {
	svc *Service
}

type seatInfo struct {
	SeatID          string `json:"seat_id"`
	UserID       	string `json:"user_id"`
	Booked 			bool   `json:"booked"`
}

func NewHandler(svc *Service) *handler {
	return &handler{
		svc,
	}
}

func (h *handler) ListSeats(w http.ResponseWriter, r *http.Request)  {
	movieID := r.PathValue("movieID")

	bookings := h.svc.ListBooking(
		movieID,
	)

	seats := make([]seatInfo, 0, len(bookings))
	for _, b := range bookings {
		seats = append(seats, seatInfo{
			SeatID: b.SeatID,
			UserID: b.UserID,
			Booked: true,
		})
	}


	utils.WriteJSON(w, http.StatusOK, seats)
}