package booking

import "sync"

type ConcurrentStore struct {
	// seats -> booking
	bookings map[string]Booking
	sync.RWMutex
}

func NewConcurrentStore() *ConcurrentStore {
	return &ConcurrentStore{
		bookings: make(map[string]Booking),
	}
}

func (s *ConcurrentStore) Book(b Booking) error {
	s.Lock()
	defer s.Unlock()

	if _, exists := s.bookings[b.SeatID]; exists {
		return ErrorSeatAlreadyBooked
	}

	s.bookings[b.SeatID] = b

	return nil
}

func (s *ConcurrentStore) ListBookings(moveID string) []Booking {
	s.RLock()
	defer s.RUnlock()

	var result []Booking

	for _, b := range s.bookings {
		if b.MovieID == moveID {
			result = append(result, b)
		}
	}

	return result
}