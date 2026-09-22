package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/qqwozz/Cinema-Booking-System/internal/booking"
	"github.com/qqwozz/Cinema-Booking-System/internal/adapters/redis"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /movies", listMovies)

	mux.Handle("GET /", http.FileServer(http.Dir("static")))

	store := booking.NewRedisStore(redis.NewClient("localhost:6379"))
	svc := booking.NewService(store)
	bookingHandler := booking.NewHandler(svc)
	mux.HandleFunc("GET /movies/{movieID}/seats", bookingHandler.ListSeats)

	mux.HandleFunc("POST /movies/{movieID}/seats/{seatID}/hold", bookingHandler.HoldSeat)

	mux.HandleFunc("PUT /sessions/{sessionID}/confirm", bookingHandler.ConfirmSession)
	mux.HandleFunc("DELETE /sessions/{sessionID}", bookingHandler.ReleaseSession)

	if err := http.ListenAndServe(":8080", mux); err != nil {
		fmt.Println("err on start web api")
		log.Fatal(err)
	}
}

func WriteJSON(w http.ResponseWriter, status int,v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

type movieResponse struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Rows        int    `json:"rows"`
	SeatsPerRow int    `json:"seats_per_row"`
}

//hardcoded movies
var movies = []movieResponse{
	{
		ID: "inception",
		Title: "Inception",
		Rows: 5,
		SeatsPerRow: 6,
	},
	{
		ID: "dune",
		Title: "Dune",
		Rows: 5,
		SeatsPerRow: 8,
	},
}

func listMovies(w http.ResponseWriter, r *http.Request) {
	WriteJSON(w, http.StatusOK, movies)
}

type seatInfo struct {
	SeatID          string `json:"seat_id"`
	UserID       	string `json:"user_id"`
	Booked 			bool   `json:"booked"`
}