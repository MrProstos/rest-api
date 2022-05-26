package server

import (
	"fmt"
	"net/http"
	"time"

	l "github.com/MrProstos/rest-api/internal/ldap"
	"github.com/dgrijalva/jwt-go/v4"
	"github.com/shaj13/go-guardian/auth"
	"github.com/shaj13/go-guardian/auth/strategies/ldap"
)

var (
	jwtKey        = []byte("my_secret_key")
	authenticator auth.Authenticator
)

type Claims struct {
	Username string
	jwt.StandardClaims
}

func readOnlyGoGuardian() {
	cfg := &ldap.Config{
		BaseDN:       "cn=API,dc=test,dc=com",
		BindDN:       "cn=read-only-admin,dc=test,dc=com",
		Port:         "389",
		Host:         "127.0.0.1",
		BindPassword: "read-admin",
		Filter:       "(cn=%s)",
	}

	authenticator = auth.New()
	strategy := ldap.New(cfg)
	authenticator.EnableStrategy(ldap.StrategyKey, strategy)

}

func Registration(w http.ResponseWriter, r *http.Request) {
	username, password, ok := r.BasicAuth()
	if len(username) == 0 && len(password) == 0 && !ok {
		http.Error(w, "fields login, password should not be empty", http.StatusBadRequest)
		return
	}

	user := l.User{Username: username, Password: password}

	err := l.ManageLDAP.AddUser(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	http.Error(w, "Authentication successful", http.StatusOK)
}

func Middleware(next http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				http.Error(w, err.Error(), http.StatusUnauthorized)
				return
			}
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		tknStr := c.Value

		claims := new(Claims)

		tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				http.Error(w, err.Error(), http.StatusUnauthorized)
				return
			}
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if !tkn.Valid {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		http.Error(w, "Successful", http.StatusOK)
	})
}

//ИСПРАВИТЬ
func Auth(w http.ResponseWriter, r *http.Request) {
	username, password, ok := r.BasicAuth()
	if len(username) == 0 && len(password) == 0 && !ok {
		http.Error(w, "fields login, password should not be empty", http.StatusBadRequest)
		return
	}
	fmt.Println(username, password)
	readOnlyGoGuardian()

	_, err := authenticator.Authenticate(r)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	expirationTime := time.Now().Add(2 * time.Hour)
	claims := &Claims{
		Username:       username,
		StandardClaims: jwt.StandardClaims{},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})
	http.Error(w, "Successful", http.StatusOK)
}
