package jwt

import (
	"crypto/ecdsa"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/rs/zerolog"
)

const ExpTimeHour = 24

type Service struct {
	privKey *ecdsa.PrivateKey
	log     zerolog.Logger
}

func NewService(privKeyPEM string, log zerolog.Logger) *Service {
	privKey, err := jwt.ParseECPrivateKeyFromPEM([]byte(privKeyPEM))
	if err != nil {
		log.Err(err).Msg("Error parsing private key")

		return nil
	}

	return &Service{
		privKey: privKey,
		log:     log,
	}
}

//nolint:errcheck
func (s *Service) New(userID string) string {
	token := jwt.New(jwt.SigningMethodES256)

	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = userID
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(time.Hour * ExpTimeHour).Unix()

	tokenString, err := token.SignedString(s.privKey)
	if err != nil {
		s.log.Err(err).Msg("Error signing token")

		return ""
	}

	return tokenString
}
