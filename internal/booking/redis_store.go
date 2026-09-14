package booking

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

const defaultHoldTTL = 7 * time.Minute

// RedisStore implements session-based seat booking backed by Redis
//
// Key design:
//
// seat:{movieid}:{seatid} -> session JSON (TTL = hold, no TTL = confirmed)
// session:{sessionID} -> seat key (reverse lookup)
type RedisStore struct {
	rdb *redis.Client
}

func NewRedisStore(rdb *redis.Client) *RedisStore {
	return &RedisStore{rdb: rdb}
}

// sessionKey builds the reverse-lookup key for a session.
func sessionKey(id string) string {
	return fmt.Sprintf("session:%s", id)
}

func (s *RedisStore) Book(b Booking) error {
	session, err := s.hold(b)
	if err != nil {
		return err
	}

	log.Printf("session booked: %v", session)

	return nil
}

func (s *RedisStore) ListBookings(moveID string) []Booking {
	return []Booking{}
}

func (s *RedisStore) hold(b Booking) (Booking, error) {
	id := uuid.New().String()
	now := time.Now()
	ctx := context.Background()
	key := fmt.Sprintf("seat:%s:%s", b.MovieID, b.SeatID)

	b.ID = id
	val, _ := json.Marshal(b)

	res := s.rdb.SetArgs(ctx, key, val, redis.SetArgs{
		Mode: "NX", // if not exists
		TTL: defaultHoldTTL,

	})
	ok := res.Val() == "OK"
	if !ok {
		return Booking{}, ErrorSeatAlreadyBooked
	}

	s.rdb.Set(ctx, sessionKey(id), key, defaultHoldTTL)

	return Booking{
		ID: id,
		MovieID: b.MovieID,
		SeatID: b.SeatID,
		UserID: b.UserID,
		Status: "held",
		ExpiresAt: now.Add(defaultHoldTTL),
	}, nil
}