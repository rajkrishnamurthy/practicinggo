package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	ldap "gopkg.in/ldap.v3"
)

var ldapHost = "129.146.97.33"
var ldapPort = "389"

var ldapUser = "rajkrishnamurthy@continube.com"

// var ldapUser = "dn: cn=ops,dc=continube,dc=com"
var ldapPassword = `L3t5G0T0TheM00n`

//
const (
	stdLongMonth      = "January"
	stdMonth          = "Jan"
	stdNumMonth       = "1"
	stdZeroMonth      = "01"
	stdLongWeekDay    = "Monday"
	stdWeekDay        = "Mon"
	stdDay            = "2"
	stdUnderDay       = "_2"
	stdZeroDay        = "02"
	stdHour           = "15"
	stdHour12         = "3"
	stdZeroHour12     = "03"
	stdMinute         = "4"
	stdZeroMinute     = "04"
	stdSecond         = "5"
	stdZeroSecond     = "05"
	stdLongYear       = "2006"
	stdYear           = "06"
	stdPM             = "PM"
	stdpm             = "pm"
	stdTZ             = "MST"
	stdISO8601TZ      = "Z0700"  // prints Z for UTC
	stdISO8601ColonTZ = "Z07:00" // prints Z for UTC
	stdNumTZ          = "-0700"  // always numeric
	stdNumShortTZ     = "-07"    // always numeric
	stdNumColonTZ     = "-07:00" // always numeric
)

var adTimeFormatInGo1 = strings.Join([]string{stdLongYear, stdZeroMonth, stdZeroDay, stdHour, stdZeroMinute, stdZeroSecond, ".0Z"}, "")

func main() {
	ldapConn, err := ldap.Dial("tcp", fmt.Sprintf("%s:%s", ldapHost, ldapPort))
	if err != nil {
		log.Fatal(err)
	}
	defer ldapConn.Close()

	// ldapConn.Start()

	err = ldapConn.Bind(ldapUser, ldapPassword)
	if err != nil {
		log.Fatal(err)
	}

	searchRequest := ldap.NewSearchRequest(
		"dc=continube,dc=com", // The base dn to search
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		"(&(objectClass=user))", // The filter to apply
		nil,
		// []string{"dn", "cn"},    // A list attributes to retrieve
		nil,
	)

	sr, err := ldapConn.Search(searchRequest)
	if err != nil {
		log.Fatal(err)
	}

	for _, entry := range sr.Entries {
		whenCreatedDTMInAD := fmt.Sprintf("%s", entry.GetAttributeValue("whenCreated"))
		// fmt.Printf("%s\n", adTimeFormatInGo1)
		fmt.Printf("%s: %v\n", entry.DN, whenCreatedDTMInAD)
		whenCreatedDTMInGo, err := time.Parse(adTimeFormatInGo1, whenCreatedDTMInAD)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Go Datetime: %s\n", whenCreatedDTMInGo)
		// entry.PrettyPrint(1)
	}

	// fmt.Printf("%s", sr)

}
