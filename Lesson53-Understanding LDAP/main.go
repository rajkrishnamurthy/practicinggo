package main

import (
	"encoding/hex"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
	"time"

	ldap "gopkg.in/ldap.v3"
)

// var ldapHost = "129.146.97.33"
var ldapHost = "activedirectory.app.continube.live"
var ldapPort = "389"

// var ldapUser = "rajkrishnamurthy@continube.com"
// var ldapUser = "dn: cn=ops,dc=continube,dc=com"
// var ldapPassword = `L3t5G0T0TheM00n`

var ldapUser = "admin-testuser@continube.com"
var ldapPassword = `$C0nt1Nub3123`

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

var adTimeFormatInGo1 = strings.Join([]string{stdLongYear, stdZeroMonth,
	stdZeroDay, stdHour,
	stdZeroMinute, stdZeroSecond, ".0Z"}, "")

//UserAccountControl :  Mapping of ?
var UserAccountControl = map[int]string{
	512:     "Enabled",
	514:     "ACCOUNTDISABLE",
	528:     " Enabled – LOCKOUT",
	530:     "ACCOUNTDISABLE – LOCKOUT",
	544:     "Enabled – PASSWD_NOTREQD",
	546:     " ACCOUNTDISABLE – PASSWD_NOTREQD",
	560:     "Enabled – PASSWD_NOTREQD – LOCKOUT",
	640:     "Enabled – ENCRYPTED_TEXT_PWD_ALLOWED",
	2048:    "INTERDOMAIN_TRUST_ACCOUNT",
	2080:    "INTERDOMAIN_TRUST_ACCOUNT – PASSWD_NOTREQD",
	4096:    "WORKSTATION_TRUST_ACCOUNT",
	8192:    " SERVER_TRUST_ACCOUNT",
	66048:   " Enabled – DONT_EXPIRE_PASSWORD",
	66050:   "ACCOUNTDISABLE – DONT_EXPIRE_PASSWORD",
	66064:   "Enabled – DONT_EXPIRE_PASSWORD – LOCKOUT",
	66066:   "ACCOUNTDISABLE – DONT_EXPIRE_PASSWORD – LOCKOUT",
	66080:   "Enabled – DONT_EXPIRE_PASSWORD – PASSWD_NOTREQD",
	66082:   "ACCOUNTDISABLE – DONT_EXPIRE_PASSWORD – PASSWD_NOTREQD",
	66176:   "Enabled – DONT_EXPIRE_PASSWORD – ENCRYPTED_TEXT_PWD_ALLOWED",
	131584:  "Enabled – MNS_LOGON_ACCOUNT",
	131586:  "ACCOUNTDISABLE – MNS_LOGON_ACCOUNT",
	131600:  "Enabled – MNS_LOGON_ACCOUNT – LOCKOUT",
	197120:  "Enabled – MNS_LOGON_ACCOUNT – DONT_EXPIRE_PASSWORD",
	532480:  "SERVER_TRUST_ACCOUNT – TRUSTED_FOR_DELEGATION (Domain Controller)",
	1049088: "Enabled – NOT_DELEGATED",
	1049090: "ACCOUNTDISABLE – NOT_DELEGATED",
	2097664: "Enabled – USE_DES_KEY_ONLY",
	2687488: " Enabled – DONT_EXPIRE_PASSWORD – TRUSTED_FOR_DELEGATION –USE_DES_KEY_ONLY",
	4194816: "Enabled – DONT_REQ_PREAUTH",
}

// https://www.epochconverter.com/ldap
// @Jan 1, 1601 UTC
var ldapZeroEpochInSeconds = int64(-11644473600)
var lastLogonInSeconds, DiffInSeconds, lastLogonHumanReadableInSeconds int64

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
		lastLogon, err := strconv.ParseInt(fmt.Sprintf("%s", entry.GetAttributeValue("lastLogon")), 10, 64)
		if err != nil || lastLogon == 0 {
			// log.Fatal(err)
			continue
		}
		lastLogonInSeconds = int64(float64(lastLogon) / math.Pow10(7))

		firstName := fmt.Sprintf("%s", entry.GetAttributeValue("givenName"))
		lastName := fmt.Sprintf("%s", entry.GetAttributeValue("sn"))
		logonName := fmt.Sprintf("%s", entry.GetAttributeValue("sAMAccountName"))
		email := fmt.Sprintf("%s", entry.GetAttributeValue("mail"))
		whenCreatedDTMInAD := fmt.Sprintf("%s", entry.GetAttributeValue("whenCreated"))
		whenCreatedDTMInGo, err := time.Parse(adTimeFormatInGo1, whenCreatedDTMInAD)

		if err != nil {
			log.Fatal(err)
		}

		whenChangedDTMInAD := fmt.Sprintf("%s", entry.GetAttributeValue("whenChanged"))
		whenChangedDTMInGo, err := time.Parse(adTimeFormatInGo1, whenChangedDTMInAD)

		if err != nil {
			log.Fatal(err)
		}
		employeeID := fmt.Sprintf("%s", entry.GetAttributeValue("employeeID"))
		employeeNumber := fmt.Sprintf("%s", entry.GetAttributeValue("employeeNumber"))
		jobTitle := fmt.Sprintf("%s", entry.GetAttributeValue("title"))
		department := fmt.Sprintf("%s", entry.GetAttributeValue("department"))
		managerName := fmt.Sprintf("%s", entry.GetAttributeValue("manager"))
		num, _ := strconv.Atoi(entry.GetAttributeValue("userAccountControl"))
		accountStatus := UserAccountControl[num]
		src := []byte(entry.GetAttributeValue("objectGUID"))
		objectGUID := hex.EncodeToString(src)
		fmt.Printf("%s", objectGUID)
		fmt.Printf("%s %s %s %s %s %s %s %s %s %s %s %s\n",
			firstName, lastName, logonName, email,
			whenCreatedDTMInGo, whenChangedDTMInGo, jobTitle, department,
			managerName, accountStatus, employeeID, employeeNumber)
		//entry.PrettyPrint(1)
		// fmt.Println()
		DiffInSeconds = lastLogonInSeconds + ldapZeroEpochInSeconds
		// lastLogonHumanReadableInSeconds = lastLogonEpochDiffInt64In100Nanos * int64(math.Pow10(-7))

		fmt.Printf("Last Logon In Seconds=%d\tDiff Epoch In Seconds =%d\tLogon Date (Human Readable) =%s\n",
			lastLogonInSeconds, DiffInSeconds, time.Unix(DiffInSeconds, 0).String())

	}

}
