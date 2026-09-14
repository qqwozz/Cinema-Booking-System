package booking

import "sync"

type MemoryStore struct {
	// seats -> booking
	bookings map[string]Booking

	sync.RWMutex
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		bookings: make(map[string]Booking),
	}
}

func (s *MemoryStore) Book(b Booking) error {
	s.Lock()
	defer s.Unlock()

	if _, exists := s.bookings[b.SeatID]; exists {
		return ErrorSeatAlreadyBooked
	}

	s.bookings[b.SeatID] = b

	return nil
}

func (s *MemoryStore) ListBookings(moveID string) []Booking {
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