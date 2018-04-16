package main

import (
	"fmt"
	"strings"

	"github.com/go-gremlin/gremlin"
)

func init() {

}

func main() {
	var querystring string

	myusername := "/dbs/graphdb/colls/Persons"
	mypassword := "Ahkrt0njQlNocrAwLKpDkJ05KzqKi1w81bHE0XoZaGy1oeed71mKJ5IpDQb36xDDAReYyybcdaXz46DxxudD0Q=="
	remotehost := "cngremlintest.gremlin.cosmosdb.azure.com"
	remoteport := "443"

	remoteurl := strings.Join([]string{"wss://", remotehost, ":", remoteport}, "")

	auth := gremlin.OptAuthUserPass(myusername, mypassword)
	client, err := gremlin.NewClient(remoteurl, auth)

	querystring = `g.V().has('firstName','Raj').bothE()`

	data, err := client.ExecQuery(querystring)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Printing Query Return \n ---------------- \n %s \n", data)
}
