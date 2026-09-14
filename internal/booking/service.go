package booking

type Service struct {
	store BookingStore
	//NOTE: add logger
}

func NewService(store BookingStore) *Service {
	return &Service{store}
}

func (s *Service) Book(b Booking) error {
	return s.store.Book(b)
}