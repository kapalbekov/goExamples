package main

import (
  "fmt"
  "strings"
  "gopkg.in/ldap.v2"
  //"crypto/tls"
)

const (
  //ldapServer = "172.16.143.51:389"
  //ldapBind = "vbtestuser"
  //ldapPassword = "Zxc123!@#"
  ldapServer = "halykbank.nb:389"
  ldapBind = "bpmuser"
  ldapPassword = "Qwe123!@#"

  filterDN = "(&(objectClass=person)(memberOf:1.2.840.113556.1.4.1941:=CN=Chat,CN=Users,DC=halykbank,DC=nb)(|(sAMAccountName={username})(mail={username})))"
  baseDN = "DC=halykbank,DC=nb"
  //baseDN = "OU=groupsOU,DC=halykbank,DC=nb"
  //baseDN = "DC=test,DC=dev"
)

func main() {
  conn, err := connect()

  if err != nil {
    fmt.Printf("Failed to connect111. %s", err)
    return
  }

  defer conn.Close()

  if err := list(conn); err != nil {
    fmt.Printf("%v___list", err)
    return
  }
}

func connect() (*ldap.Conn, error) {
  conn, err := ldap.Dial("tcp", ldapServer)

  if err != nil {
    return nil, fmt.Errorf("Failed to connect2222. %s", err)
  }
/*  
  err = conn.StartTLS(&tls.Config{InsecureSkipVerify: true})
  if err != nil {
	return nil, fmt.Errorf("Failed start TLS. %s", err)
  }
*/
  if err := conn.Bind(ldapBind, ldapPassword); err != nil {
    return nil, fmt.Errorf("Failed to bind111. %s, %s", err, "----------------")
  }

  return conn, nil
}

func list(conn *ldap.Conn) error {
  result, err := conn.Search(ldap.NewSearchRequest(
    baseDN,
    ldap.ScopeWholeSubtree,
    ldap.NeverDerefAliases,
    0,
    0,
    false,
	//"(&(objectClass=group)(!member=*))",
	"(&(objectCategory=person)(objectClass=user)(sAMAccountName=00036639))",
    //"(&(ObjectCategory=group)(cn=Департамент инновационных технологий))",
	//"(member:1.2.840.113556.1.4.1941:=(CN=00036639,CN=Users,DC=halykbank,DC=nb))",
	//"(&(objectClass=user)(mail=*test1@test.dev)(sAMAccountName=*test*))",
    []string{"dn", "sAMAccountName", "mail", "sn", "displayName", "memberOf"},
	//[]string{"dn", "memberOf"},
    nil,
  ))

  if err != nil {
    return fmt.Errorf("Failed to search users1111. %s, %s", err, "----------------")
  }

  for i, entry := range result.Entries {
    fmt.Printf(
      //"for____%s: %s %s -- %v -- %v____for\n",
      //entry.DN,
      //entry.GetAttributeValue("givenName"),
      //entry.GetAttributeValue("CN"),
      //entry.GetAttributeValue("sAMAccountName"),
      //entry.GetAttributeValue("displayName"),
	  //entry.GetAttributeValue("mail"),
	  entry.GetAttributeValue("memberOf"),
    )
	if i > 100 {
		break
	}
  }
/*  
  for i, entry := range result.Controls {
	fmt.Println("GetControlType",i, entry.GetControlType())
	fmt.Println("String",i, entry.String())
  }*/

  return nil
}



func filter(needle string) string {
  res := strings.Replace(
    filterDN,
    "{username}",
    needle,
    -1,
  )

  return res
}