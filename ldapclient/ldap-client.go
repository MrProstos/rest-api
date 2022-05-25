package ldapclient

import (
	"fmt"
	"log"

	"github.com/go-ldap/ldap/v3"
)

func AddLdapUser(username string, password string) error {
	conn, err := ldap.DialURL("ldap://127.0.0.1:389")
	if err != nil {
		return err
	}
	err = conn.Bind("cn=admin,dc=test,dc=com", "admin")
	if err != nil {
		return err
	}

	r := ldap.NewAddRequest(fmt.Sprintf("cn=%v,cn=API,dc=test,dc=com", username), []ldap.Control{})
	r.Attribute("objectClass", []string{"top", "person"})
	r.Attribute("cn", []string{username})
	r.Attribute("sn", []string{username})
	r.Attribute("userPassword", []string{password})
	err = conn.Add(r)
	if err != nil {
		return err
	}
	return nil
}

func AddToken(conn ldap.Conn) {
	r := ldap.NewModifyRequest("cn=jeck,cn=API,dc=test,dc=com", []ldap.Control{})
	r.Add("cn", []string{"token"})
	err := conn.Modify(r)
	if err != nil {
		log.Fatal(err)
	}
}

func SearchUser(conn ldap.Conn) {

	result, err := conn.Search(ldap.NewSearchRequest(
		"cn=API,dc=test,dc=com",
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		0,
		0,
		false,
		"(cn=jeck)",

		//To add anymore strings to the search, you need to add it here.
		[]string{"cn"},
		nil,
	))
	if err != nil {
		log.Fatal(err)
	}
	result.Print()
}
