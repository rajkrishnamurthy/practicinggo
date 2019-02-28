package main

import (
	"fmt"
	"os"
)

type SiteTourConfig struct {
	HelpToggleKey string // TBD
	Pages         []struct {
		Page      string // identifies the page of the site tour
		Sequences []struct {
			SeqNumber       int    // Order of flow
			ElementID       string // ID of the class in the bootstrap template file that associates with the ID
			HelpDescription string
		}
	}
}

type SiteTourRoleMapping struct {
	IncludedRoles []string // []string should be replaced w/ Role Object
	ExcludedRoles []string // []string should be replaced w/ Role Object
	*SiteTourConfig
}

func main() {

	fmt.Printf(os.Getenv("SERVICE_BUS_DNS"))

}
