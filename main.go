package main

import (
	"log"
	"net/http"

	"github.com/MrProstos/rest-api/server"
	"github.com/MrProstos/rest-api/utils"
	"github.com/gorilla/mux"
	"github.com/shaj13/go-guardian/auth"
	"github.com/shaj13/go-guardian/auth/strategies/ldap"
)

var authenticator auth.Authenticator

func main() {
	setupGoGuardian()
	utils.Logger.Info("Start server!")
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", middleware(http.HandlerFunc(server.Login))).Methods("GET")
	router.HandleFunc("/showclietns/{phone_num}", server.ShowClients).Methods("GET")
	router.HandleFunc("/addclient/", server.AddClient).Methods("POST")
	router.HandleFunc("/updateclient/", server.UpdateClient).Methods("PUT")
	router.HandleFunc("/delclient/", server.DelClient).Methods("DELETE")

	router.HandleFunc("/showorder/{client_id}", server.ShowOrder).Methods("GET")
	router.HandleFunc("/addorder/", server.AddOrder).Methods("POST")
	router.HandleFunc("/updateorder/", server.UpdateOrder).Methods("PUT")
	router.HandleFunc("/delorder/", server.DelOrder).Methods("DELETE")
	err := http.ListenAndServe(":2000", router)
	if err != nil {
		utils.Logger.Error(err.Error())
	}

}
func setupGoGuardian() {
	cfg := &ldap.Config{
		BaseDN:       "cn=Operators,dc=test,dc=com",
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

func middleware(next http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Executing Auth Middleware")
		user, err := authenticator.Authenticate(r)
		if err != nil {
			code := http.StatusUnauthorized
			http.Error(w, http.StatusText(code), code)
			return
		}
		log.Printf("User %s Authenticated\n", user.UserName())
		next.ServeHTTP(w, r)
	})
}
