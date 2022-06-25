package jwt

import (
	"fmt"
	"time"
)

type CustomClaims struct {
	Claims    map[string]interface{} `json:"data"`
	ExpiresAt int64                  `json:"expires_at"`
}

func (cc CustomClaims) Valid() error {
	if time.Now().UTC().Unix() > cc.ExpiresAt {
		return fmt.Errorf("invalid expires at")
	}

	return nil
}
