package core

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"io/ioutil"
	"time"
)

type gTokenPayload struct {
	Id   string
	Type string
}

type gTokenProvider struct {
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
}

func initTokenProvider(context Context) (tokenProvider gTokenProvider, err error) {
	privateKeyData, err := ioutil.ReadFile(context.Config.Sub("security").GetString("privateKey"))
	if err != nil {
		return
	}

	p, _ := pem.Decode(privateKeyData)
	privateKey, err := x509.ParsePKCS1PrivateKey(p.Bytes)
	if err != nil {
		return
	}

	tokenProvider = gTokenProvider{
		privateKey: privateKey,
		publicKey:  &privateKey.PublicKey,
	}

	return
}

func (t gTokenProvider) GenerateToken(payload gTokenPayload, validForHours time.Duration) (tokenString string, err error) {

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"id":   payload.Id,
		"type": payload.Type,
		"exp":  time.Now().Add(time.Hour * validForHours).Unix(),
	})
	tokenString, err = token.SignedString(t.privateKey)
	return
}

func (t gTokenProvider) VerifyToken(tokenString string) (payload gTokenPayload, err error) {
	token, err := jwt.Parse(tokenString, func(uToken *jwt.Token) (interface{}, error) {
		return t.publicKey, nil
	})

	if claim, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		payload = gTokenPayload{
			Id:   claim["id"].(string),
			Type: claim["type"].(string),
		}
	} else {
		err = errors.New("invalid token")
	}
	return
}
