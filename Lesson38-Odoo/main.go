package main

import (
	"fmt"

	"github.com/skilld-labs/go-odoo/api"
)

// https://github.com/skilld-labs/go-odoo
func main() {
	// var userIDs []int64

	url := "http://odoo.app.continube.live:8069"
	c, err := api.NewClient(url, nil)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = c.Login("Testodoo", "admin", "$C0nt1Nub3123")
	if err != nil {
		fmt.Println(err.Error())
	}

	{
		s := api.NewPurchaseOrderService(c)
		so, err := s.GetAll()
		if err != nil {
			fmt.Println(err.Error())
		}
		for _, tmpRecord := range *so {

			fmt.Printf("PO Reference:%s\n", tmpRecord.Name)
			fmt.Printf("Company Name:%s\n", tmpRecord.CompanyId.Name)
			fmt.Printf("Order Date:%s\n", tmpRecord.DateOrder)
			fmt.Printf("Vendor:%s\n", tmpRecord.PartnerId.Name)
			fmt.Printf("Amount Total:%v\n", tmpRecord.AmountTotal)
			fmt.Printf("Billing Status:%s\n", tmpRecord.InvoiceStatus.(string))

		}
	}

	// {
	// 	s := api.NewHrEmployeeService(c)
	// 	so, err := s.GetAll()
	// 	if err != nil {
	// 		fmt.Println(err.Error())
	// 	}
	// 	fmt.Println(so)
	// }

	// {
	// 	s := api.NewResGroupsService(c)
	// 	so, err := s.GetAll()
	// 	if err != nil {
	// 		fmt.Println(err.Error())
	// 	}
	// 	for _, tmpGrp := range *so {

	// 		if strings.Contains(strings.ToLower(tmpGrp.DisplayName), "accounting") {
	// 			// fmt.Printf("Name = %v \n", tmpGrp.Name)
	// 			fmt.Printf("Group Name = %v \n", tmpGrp.DisplayName)
	// 			userIDs = tmpGrp.Users
	// 			fmt.Printf("Users = %#v \n", tmpGrp.Users)
	// 			fmt.Printf("ModelAccess = %#v \n", tmpGrp.ModelAccess)
	// 		}

	// 	}
	// }

	// {
	// 	s := api.NewIrModelService(c)
	// 	so, err := s.GetByIds([]int64{190, 194, 195, 197, 200, 209, 210, 222, 239, 243, 246, 248, 250, 252, 254, 256, 257, 259, 261, 264, 265})
	// 	if err != nil {
	// 		fmt.Println(err.Error())
	// 	}
	// 	for _, tmpUser := range *so {
	// 		// if tmpUser.Name == "Administrator" {
	// 		// 	fmt.Printf("%#v \n", tmpUser)
	// 		// }
	// 		fmt.Printf("Model Name = %v \n", tmpUser.DisplayName)
	// 		// fmt.Printf("Name = %v \n", tmpUser.Name)
	// 	}
	// }

	// {
	// 	s := api.NewResUsersService(c)
	// 	so, err := s.GetByIds(userIDs)
	// 	if err != nil {
	// 		fmt.Println(err.Error())
	// 	}
	// 	for _, tmpUser := range *so {
	// 		// if tmpUser.Name == "Administrator" {
	// 		// 	fmt.Printf("%#v \n", tmpUser)
	// 		// }
	// 		fmt.Printf("User Name = %v \n", tmpUser.DisplayName)
	// 		// fmt.Printf("Name = %v \n", tmpUser.Name)
	// 	}
	// }
}
