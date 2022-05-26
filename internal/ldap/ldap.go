package ldap

import (
	"fmt"
	"github.com/go-ldap/ldap/v3"
	"log"
)

type ManageLDAP interface {
	AddUser() (err error)
}

type User struct {
	Username string
	Password string
}

func (user User) AddUser() (err error) {
	conn := GetLDAP()

	request := ldap.NewAddRequest(fmt.Sprintf("cn=%v,cn=API,dc=test,dc=com", user.Username), []ldap.Control{})
	request.Attribute("objectClass", []string{"top", "person"})
	request.Attribute("cn", []string{user.Username})
	request.Attribute("sn", []string{user.Username})
	request.Attribute("userPassword", []string{user.Password})

	err = conn.Add(request)
	if err != nil {
		return
	}
	return
}

func SearchUser(username string, password string) {
	conn := GetLDAP()
	result, err := conn.Search(ldap.NewSearchRequest(
		"cn=API,dc=test,dc=com",
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		0,
		0,
		false,
		fmt.Sprintf("(cn=%v)", username),
		[]string{"cn"},
		nil,
	))
	if err != nil {
		log.Fatalln(err)
	}
	result.Print()
}
