package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/MrProstos/rest-api/db"
	"github.com/MrProstos/rest-api/ldapclient"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
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

		fmt.Fprintf(w, "Succses!")
	})
}

func Registration(w http.ResponseWriter, r *http.Request) {
	username, password, ok := r.BasicAuth()
	if len(username) == 0 && len(password) == 0 && !ok {
		http.Error(w, "fields login, password should not be empty", http.StatusBadRequest)
		return
	}
	err := ldapclient.AddLdapUser(username, password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Fprint(w, "Authentication successful")
}

func Auth(w http.ResponseWriter, r *http.Request) {
	username, password, ok := r.BasicAuth()
	if len(username) == 0 && len(password) == 0 && !ok {
		http.Error(w, "fields login, password should not be empty", http.StatusBadRequest)
		return
	}
	fmt.Println(username, password)
	readOnlyGoGuardian()

	info, err := authenticator.Authenticate(r)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}
	fmt.Println(info)

	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println(tokenString)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})
	fmt.Fprint(w, http.StatusOK)
}

func ShowClients(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	client := new(db.Client)
	client.Phone_num = vars["phone_num"]

	err := db.Db_manage.Show(client)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "%v %v", http.StatusOK, client)
}

func AddClient(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	client := new(db.Client)
	err := decoder.Decode(&client)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = db.Db_manage.Add(client)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprint(w, http.StatusOK)
}

func UpdateClient(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	client := new(db.Client)
	if err := decoder.Decode(&client); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := db.Db_manage.Update(client)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprint(w, http.StatusOK)

}

func DelClient(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	client := new(db.Client)
	err := decoder.Decode(&client)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = db.Db_manage.Del(client)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprint(w, http.StatusOK)
}

func ShowOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	ord := new(db.Order)
	str_client_id := vars["client_id"]

	int_client_id, err := strconv.Atoi(str_client_id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	ord.Client_id = uint(int_client_id)

	err = db.Db_manage.Show(ord)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "%v %v", http.StatusOK, ord)

}

func AddOrder(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	order := new(db.Order)
	err := decoder.Decode(&order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = db.Db_manage.Add(order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprint(w, http.StatusOK)
}

func UpdateOrder(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	order := new(db.Order)
	err := decoder.Decode(&order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = db.Db_manage.Update(order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprint(w, http.StatusOK)

}

func DelOrder(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	order := new(db.Order)
	err := decoder.Decode(&order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = db.Db_manage.Del(order)
	if err != nil {

		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprint(w, http.StatusOK)
}
