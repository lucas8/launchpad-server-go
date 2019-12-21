package jwt

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"

	"github.com/go-chi/jwtauth"
)

type Token struct {
	Claims    Claims
	tokenAuth *jwtauth.JWTAuth
}

type Claims struct {
	jwt.StandardClaims
	UserID uint   `json:"user_id"`
	Role   string `json:"role"`
}

// contextKey is a value for use with context.WithValue. It's used as
// a pointer so it fits in an interface{} without allocation. This technique
// for defining context keys was copied from Go 1.7's new use of context in net/http.
type contextKey struct {
	name string
}

var TokenCtxKey = &contextKey{"Token"}

// Creates new jwt with valid jwtauth instance
func (Token) New() *Token {
	token := &Token{
		tokenAuth: jwtauth.New("HS256", []byte(os.Getenv("ACCESS_TOKEN_SECRET")), nil),
	}

	return token
}

func (t *Token) Encode(UserID uint, Role string) string {
	claims := &Claims{
		UserID: UserID,
		Role:   Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "getlaunchpad.dev",
		},
	}

	_, tokenString, err := t.tokenAuth.Encode(claims)
	if err != nil {
		log.Panic(err)
		return ""
	}

	return tokenString
}

func (t *Token) Verifier() func(http.Handler) http.Handler {
	return jwtauth.Verifier(t.tokenAuth)
}

func (t *Token) Decode(r *http.Request) *Claims {
	val, _ := t.Authenticate(r)
	return val
}

func (t *Token) Authenticate(r *http.Request) (*Claims, error) {
	token, _, err := jwtauth.FromContext(r.Context())
	if err != nil {
		return &Claims{}, errors.New("Invalid JWT")
	}

	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(token.Raw, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("ACCESS_TOKEN_SECRET")), nil
	})
	if err != nil || tkn == nil {
		return &Claims{}, errors.Wrap(err, "Empty or invalid JWT")
	}
	if !tkn.Valid {
		return &Claims{}, errors.New("Invalid JWT")
	}

	return claims, nil
}

func (t *Token) Authenticator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := t.Authenticate(r)

		if err != nil {
			fmt.Println(err)
			http.Error(w, http.StatusText(401), 401)
			return
		}

		// Token is authenticated, pass it through
		next.ServeHTTP(w, r)
	})
}
