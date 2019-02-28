package main

import (
	"crypto"
	"fmt"

	"github.com/sec51/twofactor"
)

func main() {

	otp, err := twofactor.NewTOTP("raj.krishnamurthy@continube.com", "ContiNube", crypto.SHA1, 8)
	if err != nil {
		fmt.Printf("%v \n", err)
	}

	qrBytes, err := otp.QR()
	if err != nil {
		fmt.Printf("%v \n", err)
	}

	fmt.Printf("qrBytes\n---------\n%v\n", qrBytes)

}
