package jwt

import (
	"fmt"

	"github.com/arthureichelberger/trailrcore/pkg/env"
	"github.com/golang-jwt/jwt"
	"github.com/rs/zerolog/log"
)

var jwtSecret string = env.Get("TRAILRCORE_JWT_SECRET", "verysecret")

func New(claims CustomClaims) (string, error) {
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := jwtToken.SignedString([]byte(jwtSecret))
	if err != nil {
		log.Error().Err(err).Interface("claims", claims).Msg("could not sign jwt")
		return "", fmt.Errorf("could not sign jwt")
	}

	return token, nil
}
