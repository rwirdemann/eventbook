package rest

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"os"
	"regexp"
)

func JWTAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := extractJwtFromHeader(r.Header)
		if verifyJWT(token) {
			next.ServeHTTP(w, r)
			return
		}
		w.Header().Set("WWW-Authenticate", "Bearer realm.go=\"eventbook.de\"")
		w.WriteHeader(http.StatusUnauthorized)
	}
}

func extractJwtFromHeader(header http.Header) (jwt string) {
	var jwtRegex = regexp.MustCompile(`^Bearer (\S+)$`)

	if val, ok := header["Authorization"]; ok {
		for _, value := range val {
			if result := jwtRegex.FindStringSubmatch(value); result != nil {
				jwt = result[1]
				return
			}
		}
	}

	return
}

func verifyJWT(token string) bool {
	secretKey := "-----BEGIN CERTIFICATE-----\n" + os.Getenv("KEYCLOAK_RS256_PUB_KEY") + "\n-----END CERTIFICATE-----"
	key, err := jwt.ParseRSAPublicKeyFromPEM([]byte(secretKey))
	if err != nil {
		return false
	}

	t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return key, nil
	})

	return err == nil && t.Valid
}
