package token

import (
	"encoding/json"
	"time"

	"aidanwoods.dev/go-paseto"
)

type PasetoMaker struct {
	parser       paseto.Parser
	symmetricKey paseto.V4SymmetricKey
}

// CreateToken implements Maker.
func (maker *PasetoMaker) CreateToken(name string, duration time.Duration) (string, error) {
	payload, err := NewPayload(name, duration)
	if err != nil {
		return "", err
	}

	token := paseto.NewToken()
	token.SetIssuedAt(payload.IssuedAt)
	token.SetNotBefore(time.Now())
	token.SetExpiration(payload.ExpiredAt)

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	token.SetString("payload", string(jsonData))

	return token.V4Encrypt(maker.symmetricKey, nil), nil
}

var payload Payload

// VerifyToken implements Maker.
func (maker *PasetoMaker) VerifyToken(token string) (*Payload, error) {
	parsedToken, err := maker.parser.ParseV4Local(maker.symmetricKey, token, nil)
	if err != nil {
		return nil, err
	}

	p, err := parsedToken.GetString("payload")
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(p), &payload)
	if err != nil {
		return nil, err
	}

	return &payload, nil
}

func NewPasetoMaker(secret string) (*PasetoMaker, error) {
	secretKey, err := paseto.V4SymmetricKeyFromHex(secret)
	if err != nil {
		return nil, err
	}

	parser := paseto.NewParser()
	parser.AddRule(paseto.NotExpired())

	maker := &PasetoMaker{
		parser:       parser,
		symmetricKey: secretKey,
	}
	return maker, nil
}
