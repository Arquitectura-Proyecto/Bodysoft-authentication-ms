package data

import (
	"crypto/tls"
	"fmt"

	"gopkg.in/ldap.v3"
)

//LDAPConnection ..
func LDAPConnection() *ldap.Conn {
	bindusername := "cn=admin,dc=bodysoft,dc=unal,dc=edu,dc=co"
	bindpassword := "admin"
	l, err := ldap.Dial("tcp", fmt.Sprintf("%s:%d", "192.168.0.3", 389))
	if err != nil {
		panic(err.Error())
	}
	err = l.StartTLS(&tls.Config{InsecureSkipVerify: true})
	if err != nil {
		panic(err.Error())
	}
	err = l.Bind(bindusername, bindpassword)
	if err != nil {
		panic(err.Error())
	}
	return l
}

//LDAPSearchRequest ..
func LDAPSearchRequest(username string) *ldap.SearchRequest {
	bindusername := "cn=admin,dc=bodysoft,dc=unal,dc=edu,dc=co"
	bindpassword := "admin"
	l, err := ldap.Dial("tcp", fmt.Sprintf("%s:%d", "192.168.0.3", 389))
	if err != nil {
		panic(err.Error())
	}
	err = l.StartTLS(&tls.Config{InsecureSkipVerify: true})
	if err != nil {
		panic(err.Error())
	}
	err = l.Bind(bindusername, bindpassword)
	if err != nil {
		panic(err.Error())
	}
	searchRequest := ldap.NewSearchRequest(
		"dc=bodysoft,dc=unal,dc=edu,dc=co",
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		fmt.Sprintf("(&(objectClass=organizationalPerson)(cn=%s))", username),
		[]string{"dn description"},
		nil,
	)
	return searchRequest
}
