package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	gourl "net/url"
	"strings"
	"time"
)

func main() {

	var accessToken string
	url := "https://login.microsoftonline.com/6fb02b11-3836-4731-aace-658d73bf9ac4/oauth2/v2.0/token"
	form := gourl.Values{}
	form.Add("grant_type", "client_credentials")
	form.Add("client_id", "6e84e074-1598-4904-8d20-2adab2db966e")
	form.Add("client_secret", "3UCHWGkpW4cQ4zOqbYD?kZrtcbXU.8[*")
	form.Add("scope", "https://graph.microsoft.com/.default")
	req, err := http.NewRequest("POST", url, strings.NewReader(form.Encode()))
	if err != nil {
		log.Fatalf("%s", err)
	}
	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("%s", err)
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	var jsondata ClientCredential
	err = json.Unmarshal([]byte(body), &jsondata)
	if err != nil {
		log.Fatalf("%s", err)
	}
	accessToken = "Bearer " + jsondata.Access_token

	users, err := getUsers(accessToken)
	if err != nil {
		log.Fatalf("%s", err)
	}

	for _, user := range users {
		// getUserAsMemberOf(accessToken, user)
		getUserDetails(accessToken, user)

	}

}

func getUsers(accessToken string) (users map[string]string, err error) {
	microsoftURL := "https://graph.microsoft.com/beta/users"
	req, err := http.NewRequest("GET", microsoftURL, nil)
	if err != nil {
		log.Fatalf("%s", err)
	}
	req.Header.Add("Authorization", accessToken)
	res, _ := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
		// log.Fatalf("%s", err)

	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	// fmt.Printf("Printing output from /users GET call\n%s\n", body)

	// var userDetails Users
	var userDetails ODataUsers

	err = json.Unmarshal([]byte(body), &userDetails)
	if err != nil {
		return nil, err
		// log.Fatalf("%s", err)
	}

	// azureADList := make([]AzureADUserData, 0)
	users = make(map[string]string)
	for _, userRecord := range userDetails.Value {
		// createdDateTime := GetGoDateFormat(userDetails.Users[m].CreatedDateTime)
		// data := AzureADUserData{
		// 	GivenName:       userDetails.Users[m].GivenName,
		// 	Surname:         userDetails.Users[m].Surname,
		// 	DisplayName:     userDetails.Users[m].DisplayName,
		// 	Mail:            userDetails.Users[m].Mail,
		// 	CreatedDateTime: createdDateTime,
		// 	EmployeeId:      userDetails.Users[m].EmployeeId,
		// 	JobTitle:        userDetails.Users[m].JobTitle,
		// 	Department:      userDetails.Users[m].Department,
		// 	AccountEnabled:  userDetails.Users[m].AccountEnabled,
		// 	Id:              userDetails.Users[m].Id,
		// }
		// azureADList = append(azureADList, data)
		users[userRecord.MailNickname] = userRecord.ID
	}
	// fileName := fmt.Sprintf("%v-%v%v", "AzureADUserList", time.Now().Unix(), ".json")
	// file, _ := json.MarshalIndent(azureADList, "", "\t")
	// _ = ioutil.WriteFile(fileName, file, 0644)

	return users, nil
}

func getUserDetails(accessToken string, userID string) (err error) {
	microsoftURL := fmt.Sprintf("%s/%s", "https://graph.microsoft.com/beta/users", userID)
	req, err := http.NewRequest("GET", microsoftURL, nil)
	if err != nil {
		log.Fatalf("%s", err)
	}
	req.Header.Add("Authorization", accessToken)
	res, _ := http.DefaultClient.Do(req)
	if err != nil {
		return err
		// log.Fatalf("%s", err)

	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Printf("Printing output from /users GET call\n%s\n---------\n", body)

	// // var userDetails Users
	// var userDetails ODataUsers

	// err = json.Unmarshal([]byte(body), &userDetails)
	// if err != nil {
	// 	return err
	// 	// log.Fatalf("%s", err)
	// }

	return nil
}

func getUserAsMemberOf(accessToken string, userID string) (err error) {
	microsoftURL := fmt.Sprintf("%s/%s/memberOf", "https://graph.microsoft.com/beta/users", userID)
	req, err := http.NewRequest("GET", microsoftURL, nil)
	if err != nil {
		log.Fatalf("%s", err)
	}
	req.Header.Add("Authorization", accessToken)
	res, _ := http.DefaultClient.Do(req)
	if err != nil {
		return err
		// log.Fatalf("%s", err)

	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Printf("Printing output from /users GET call\n%s\n---------\n", body)

	// // var userDetails Users
	// var userDetails ODataUsers

	// err = json.Unmarshal([]byte(body), &userDetails)
	// if err != nil {
	// 	return err
	// 	// log.Fatalf("%s", err)
	// }

	return nil
}

func GetGoDateFormat(dateTMInAD string) string {
	var dateTMInGo string
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
	var adTimeFormatInGo = strings.Join([]string{stdLongYear, "-", stdZeroMonth, "-", stdZeroDay, "T", stdHour, ":", stdZeroMinute, ":", stdZeroSecond, "Z"}, "")
	DateTMInGo, err := time.Parse(adTimeFormatInGo, dateTMInAD)
	if err != nil {
		log.Fatal(err)
	}
	dateTMInGo = DateTMInGo.Format("2006-Jan-2")
	return dateTMInGo
}

type ClientCredential struct {
	Access_token string `json:"access_token"`
}

type AzureADUserData struct {
	GivenName       string
	Surname         string
	DisplayName     string
	Mail            string
	CreatedDateTime string
	EmployeeId      string
	JobTitle        string
	Department      string
	AccountEnabled  bool
	Id              string
}
type Users struct {
	Users []AzureADUserData `json:"value"`
}

// ODataUsers : GET /users response
type ODataUsers struct {
	OdataContext string `json:"@odata.context"`
	Value        []struct {
		ID                              string        `json:"id"`
		DeletedDateTime                 interface{}   `json:"deletedDateTime"`
		AccountEnabled                  bool          `json:"accountEnabled"`
		AgeGroup                        interface{}   `json:"ageGroup"`
		BusinessPhones                  []interface{} `json:"businessPhones"`
		City                            interface{}   `json:"city"`
		CreatedDateTime                 time.Time     `json:"createdDateTime"`
		CompanyName                     interface{}   `json:"companyName"`
		ConsentProvidedForMinor         interface{}   `json:"consentProvidedForMinor"`
		Country                         interface{}   `json:"country"`
		Department                      interface{}   `json:"department"`
		DisplayName                     string        `json:"displayName"`
		EmployeeID                      interface{}   `json:"employeeId"`
		FaxNumber                       interface{}   `json:"faxNumber"`
		GivenName                       interface{}   `json:"givenName"`
		ImAddresses                     []interface{} `json:"imAddresses"`
		IsResourceAccount               interface{}   `json:"isResourceAccount"`
		JobTitle                        interface{}   `json:"jobTitle"`
		LegalAgeGroupClassification     interface{}   `json:"legalAgeGroupClassification"`
		Mail                            interface{}   `json:"mail"`
		MailNickname                    string        `json:"mailNickname"`
		MobilePhone                     interface{}   `json:"mobilePhone"`
		OnPremisesDistinguishedName     interface{}   `json:"onPremisesDistinguishedName"`
		OfficeLocation                  interface{}   `json:"officeLocation"`
		OnPremisesDomainName            interface{}   `json:"onPremisesDomainName"`
		OnPremisesImmutableID           interface{}   `json:"onPremisesImmutableId"`
		OnPremisesLastSyncDateTime      interface{}   `json:"onPremisesLastSyncDateTime"`
		OnPremisesSecurityIdentifier    interface{}   `json:"onPremisesSecurityIdentifier"`
		OnPremisesSamAccountName        interface{}   `json:"onPremisesSamAccountName"`
		OnPremisesSyncEnabled           interface{}   `json:"onPremisesSyncEnabled"`
		OnPremisesUserPrincipalName     interface{}   `json:"onPremisesUserPrincipalName"`
		OtherMails                      []interface{} `json:"otherMails"`
		PasswordPolicies                interface{}   `json:"passwordPolicies"`
		PostalCode                      interface{}   `json:"postalCode"`
		PreferredDataLocation           interface{}   `json:"preferredDataLocation"`
		PreferredLanguage               interface{}   `json:"preferredLanguage"`
		ProxyAddresses                  []interface{} `json:"proxyAddresses"`
		RefreshTokensValidFromDateTime  time.Time     `json:"refreshTokensValidFromDateTime"`
		ShowInAddressList               interface{}   `json:"showInAddressList"`
		SignInSessionsValidFromDateTime time.Time     `json:"signInSessionsValidFromDateTime"`
		State                           interface{}   `json:"state"`
		StreetAddress                   interface{}   `json:"streetAddress"`
		Surname                         interface{}   `json:"surname"`
		UsageLocation                   interface{}   `json:"usageLocation"`
		UserPrincipalName               string        `json:"userPrincipalName"`
		ExternalUserState               interface{}   `json:"externalUserState"`
		ExternalUserStateChangeDateTime interface{}   `json:"externalUserStateChangeDateTime"`
		UserType                        string        `json:"userType"`
		AssignedLicenses                []interface{} `json:"assignedLicenses"`
		AssignedPlans                   []interface{} `json:"assignedPlans"`
		DeviceKeys                      []interface{} `json:"deviceKeys"`
		OnPremisesExtensionAttributes   struct {
			ExtensionAttribute1  interface{} `json:"extensionAttribute1"`
			ExtensionAttribute2  interface{} `json:"extensionAttribute2"`
			ExtensionAttribute3  interface{} `json:"extensionAttribute3"`
			ExtensionAttribute4  interface{} `json:"extensionAttribute4"`
			ExtensionAttribute5  interface{} `json:"extensionAttribute5"`
			ExtensionAttribute6  interface{} `json:"extensionAttribute6"`
			ExtensionAttribute7  interface{} `json:"extensionAttribute7"`
			ExtensionAttribute8  interface{} `json:"extensionAttribute8"`
			ExtensionAttribute9  interface{} `json:"extensionAttribute9"`
			ExtensionAttribute10 interface{} `json:"extensionAttribute10"`
			ExtensionAttribute11 interface{} `json:"extensionAttribute11"`
			ExtensionAttribute12 interface{} `json:"extensionAttribute12"`
			ExtensionAttribute13 interface{} `json:"extensionAttribute13"`
			ExtensionAttribute14 interface{} `json:"extensionAttribute14"`
			ExtensionAttribute15 interface{} `json:"extensionAttribute15"`
		} `json:"onPremisesExtensionAttributes"`
		OnPremisesProvisioningErrors []interface{} `json:"onPremisesProvisioningErrors"`
		PasswordProfile              struct {
			Password                             interface{} `json:"password"`
			ForceChangePasswordNextSignIn        bool        `json:"forceChangePasswordNextSignIn"`
			ForceChangePasswordNextSignInWithMfa bool        `json:"forceChangePasswordNextSignInWithMfa"`
		} `json:"passwordProfile"`
		ProvisionedPlans []interface{} `json:"provisionedPlans"`
	} `json:"value"`
}
