package token

import "time"

type Maker interface {
	CreateToken(userID int64, duration time.Duration) (string, *Payload, error)
	VerifyToken(token string) (*Payload, error)
}
