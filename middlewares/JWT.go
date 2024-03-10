package middlewares

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
)

var JwtKey = []byte("secret-key")

type Claims struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.StandardClaims
}

func CreateJWT(id int, user string, email string, w http.ResponseWriter, r *http.Request) {

	expirationTime := time.Now().Add(time.Minute * 1000)

	claims := &Claims{
		Id:       id,
		Username: user,
		Email:    email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(JwtKey)

	if err != nil {
		return
	}

	http.SetCookie(w,
		&http.Cookie{
			Name:     "token",
			Value:    tokenString,
			HttpOnly: true,
			Secure:   true,
			Path:     "/",
			Expires:  expirationTime,
		})
}

func VerifyJWT(w http.ResponseWriter, r *http.Request) (bool, string, int) {
	cookie, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			return false, "", 0
		}
	}

	tokenStr := cookie.Value

	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(tokenStr, claims,
		func(t *jwt.Token) (interface{}, error) {
			return JwtKey, nil
		})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return false, "", 0
		}
		return false, "", 0
	}

	if !tkn.Valid {
		return false, "", 0
	}

	return true, string(claims.Username), claims.Id
}
