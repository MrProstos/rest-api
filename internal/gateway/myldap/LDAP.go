package myldap

import (
	"fmt"
	"github.com/go-ldap/ldap/v3"
	"log"
)

type operator struct {
	Username string
	Password string
}

func NewOperator(username string, password string) *operator {
	return &operator{Username: username, Password: password}
}

func (operator *operator) AddUser() error {
	request := ldap.NewAddRequest(fmt.Sprintf("cn=%v,cn=API,dc=test,dc=com", operator.Username), []ldap.Control{})
	request.Attribute("objectClass", []string{"top", "person"})
	request.Attribute("cn", []string{operator.Username})
	request.Attribute("sn", []string{operator.Username})
	request.Attribute("userPassword", []string{operator.Password})

	err := getLDAP().Add(request)
	if err != nil {
		return err
	}
	return nil
}

func (operator *operator) Search() error {
	result, err := getLDAP().Search(ldap.NewSearchRequest(
		"cn=API,dc=test,dc=com",
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		0,
		0,
		false,
		fmt.Sprintf("(cn=%v)", operator.Username),
		[]string{"cn"},
		nil,
	))
	if err != nil {
		log.Fatalln(err)
	}
	result.PrettyPrint(1)
	return nil
}
