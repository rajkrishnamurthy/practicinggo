package main

import (
	"cncommonlibs/constants"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/lib/pq"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	testJSONOps()
}

func testJSONOps() {
	jsonOutput := make([]*AppCred, 0)
	{
		s1 := &AppCred{}
		s1.AppGUID = "0ebec76c-20ff-402f-bbb9-fca92e3c93e9"
		s1.Creds = make([]*Credential, 0)
		{
			tmp := &Credential{}
			tmp.UserID = "john.doe@acme.com"
			tmp.UserDomain = "acme.continube.com"
			tmp.SSHPrivateKey = []byte(sshkey)
			tmp.OtherCredInfo = make(map[string]interface{})
			tmp.OtherCredInfo["database"] = "TestOdoo"
			s1.Creds = append(s1.Creds, tmp)
		}

		{
			tmp := &Credential{}
			tmp.UserID = "jane.doe@fabrikam.com"
			tmp.UserDomain = "fabrikam.continube.com"
			tmp.SSHPrivateKey = []byte(sshkey)
			tmp.OtherCredInfo = make(map[string]interface{})
			tmp.OtherCredInfo["manager"] = "jen.doe@fabrikam.com"
			s1.Creds = append(s1.Creds, tmp)

		}
		jsonOutput = append(jsonOutput, s1)
	}
	{
		s1 := &AppCred{}
		s1.AppGUID = "503846da-71cc-47c8-aad3-9335880f60f5"
		s1.Creds = make([]*Credential, 0)
		{
			tmp := &Credential{}
			tmp.UserID = "loga.vinayagam@continube.com"
			tmp.UserDomain = "continube.com"
			tmp.SSHPrivateKey = []byte(sshkey)
			tmp.OtherCredInfo = make(map[string]interface{})
			tmp.OtherCredInfo["database"] = "SAP"
			s1.Creds = append(s1.Creds, tmp)
		}

		{
			tmp := &Credential{}
			tmp.UserID = "jane.doe@fabrikam.com"
			tmp.UserDomain = "fabrikam.continube.com"
			tmp.SSHPrivateKey = []byte(sshkey)
			tmp.OtherCredInfo = make(map[string]interface{})
			tmp.OtherCredInfo["manager"] = "jen.doe@fabrikam.com"
			s1.Creds = append(s1.Creds, tmp)

		}
		jsonOutput = append(jsonOutput, s1)
	}

	payload, err := json.Marshal(jsonOutput)
	if err != nil {
		fmt.Printf("Cannot marshal %v \n", err)
	}

	fmt.Printf("%s", payload)

}

func testDBOps() {
	constants.InitConstants()
	connectionstring := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		constants.CNdbuser, constants.CNdbpassword,
		constants.CNdbhostname, constants.CNdbport,
		constants.CNdbname, "disable")

	db, err := gorm.Open("postgres", connectionstring)
	if err != nil {
		log.Printf("gorm.Open() Error = %v \n", err)
	}

	// db.DropTableIfExists(AppGroupModel{})

	if !db.HasTable(AppGroupModel{}) {
		err := db.AutoMigrate(AppGroupModel{}).Error
		if err != nil {
			log.Printf("Error in AutoMigrate(AppGroupModel{}): %v", err)
		}
	}
	appgrp1 := &AppGroupModel{}
	appgrp1.ID = 10
	appgrp1.CreatedAt, appgrp1.UpdatedAt = time.Now(), time.Now()
	appgrp1.AppGroupType = 2
	appgrp1.Applications = pq.Int64Array{1, 2, 3, 4}
	appgrp1.UserDomain = "dummy.dummy.com"
	appGroupTags := make(map[string]string)
	appGroupTags["app"] = "odoo"
	appGroupTags["os"] = "linux"
	tmp, _ := json.Marshal(appGroupTags)
	appgrp1.AppTags = tmp //fmt.Sprintf("%s", tmp)

	fmt.Printf("%v", db.Save(appgrp1).Error)
}
