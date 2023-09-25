package token

import (
	"time"
)

type Payload struct {
	Name      string    `json:"name"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

func NewPayload(name string, duration time.Duration) (*Payload, error) {
	payload := &Payload{
		Name:      name,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}

	return payload, nil
}
