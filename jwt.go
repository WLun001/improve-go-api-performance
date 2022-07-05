package main

import (
	"gopkg.in/square/go-jose.v2"
	"gopkg.in/square/go-jose.v2/jwt"
	"time"
)

var jwtSecret = []byte("super-secret")

type JWTToken struct {
	Token     string
	IssuedAt  time.Time
	ExpiresAt time.Time
}

type Claims struct {
	*jwt.Claims
	Secret string `json:"secret,omitempty"`
}

func createJWTToken(user, payload string) (*JWTToken, error) {
	issuedAt := time.Now().UTC()
	expiresAt := issuedAt.Add(time.Minute * 15)
	claims := Claims{
		Claims: &jwt.Claims{
			Subject:  user,
			Expiry:   jwt.NewNumericDate(expiresAt),
			IssuedAt: jwt.NewNumericDate(issuedAt),
		},
		Secret: payload,
	}

	var signerOpts = jose.SignerOptions{}
	signerOpts.WithType("JWT")
	sign, err := jose.NewSigner(
		jose.SigningKey{
			Algorithm: jose.HS256,
			Key:       jwtSecret,
		}, &signerOpts,
	)
	if err != nil {
		return nil, err
	}

	token, err := jwt.Signed(sign).Claims(claims).CompactSerialize()
	if err != nil {
		return nil, err
	}
	return &JWTToken{
		Token:     token,
		IssuedAt:  issuedAt,
		ExpiresAt: expiresAt,
	}, nil

}

func verifyJWTToken(token string) (*Claims, error) {
	enc, err := jwt.ParseSigned(token)
	if err != nil {
		return nil, err
	}

	claim := new(Claims)
	if err := enc.Claims(jwtSecret, claim); err != nil {
		return nil, err
	}

	if claim.Expiry.Time().UTC().Unix()-time.Now().UTC().Unix() < 0 {
		return claim, jwt.ErrExpired
	}
	return claim, nil
}
