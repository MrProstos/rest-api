package main

import (
	"fmt"
	"log"

	"github.com/go-ldap/ldap"
)

const (
	Username = "vmikhin"
	Password = "Zz123456"
	base     = "cn=admin,dc=example,dc=org"
	filter   = "(&(objectClass=*))"
)

func Connect() (*ldap.Conn, error) {
	l, err := ldap.Dial("tcp", fmt.Sprintf("%s:%d", "example.org", 389))
	if err != nil {
		return nil, err
	}
	return l, nil

}

func main() {
	_, err := Connect()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("NIL")
}
