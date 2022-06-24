package myldap

import (
	"github.com/go-ldap/ldap/v3"
	"log"
)

const (
	url      string = "ldap://127.0.0.1:389"
	bind     string = "cn=admin,dc=test,dc=com"
	password string = "admin"
)

var conn *ldap.Conn

type conf struct {
	url      string
	bind     string
	password string
}

func newConf(url string, bind string, password string) *conf {
	return &conf{url: url, bind: bind, password: password}
}

func (conf *conf) Connect() error {
	dial, err := ldap.DialURL(conf.url)
	if err != nil {
		log.Fatalln(err)
	}

	err = dial.Bind(conf.bind, conf.password)
	if err != nil {
		log.Fatalln(err)
	}
	conn = dial
	return nil
}

func GetLDAP() *ldap.Conn {
	return conn
}
