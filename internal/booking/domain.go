package booking

import "errors"

var (
	ErrorSeatAlreadyBooked = errors.New("Seat is already taken")
)

type Booking struct {
	ID       string
	MovieID  string
	SeatID   string
	UserID   string
	StatusID string
}

type BookingStore interface {
	Book(b Booking) error
	ListBookings(movieId string) []Booking
}