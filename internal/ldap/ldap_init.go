package ldap

import (
	"github.com/go-ldap/ldap/v3"
	"log"
)

var Conn *ldap.Conn

const (
	url      string = "ldap://127.0.0.1:389"
	bind     string = "cn=admin,dc=test,dc=com"
	password string = "admin"
)

func init() {
	conn, err := ldap.DialURL(url)
	if err != nil {
		log.Fatalln(err)
	}

	err = conn.Bind(bind, password)
	if err != nil {
		log.Fatalln(err)
	}

	Conn = conn
}

// GetLDAP возвращает дескриптор объекта Conn
func GetLDAP() *ldap.Conn {
	return Conn
}
