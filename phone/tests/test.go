package main

import (
	"fmt"
	"github.com/nkozyra/entities/phone"
)

func main() {


	thisPhone := phone.New("1 (555)  123     4567")
	thisPhone.Normalize()
	fmt.Println(thisPhone)

	thatPhone := phone.New("+1 813 598.4742")
	thatPhone.Normalize()
	fmt.Println(thatPhone)
}
