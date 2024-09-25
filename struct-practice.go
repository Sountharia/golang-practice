// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	type wheel struct {
		radius int
		psi    int
	}
	
	// NESTED STRUCT
	type car struct {
		make       string
		model      string
		frontwheel wheel
	}

	// EMBEDDED STRUCT
	type taxi struct {
		car
		owner string
	}

	// INTIALIZING NESTED STRUCT
	myCar := car{}
	myCar.make = "test"
	myCar.model = "mod"
	myCar.frontwheel.radius = 10

	// INTIALIZING EMBEDDED STRUCT
	myTaxi := taxi{
		owner: "sou",
		car: car{
			make:  "toyota",
			model: "camry",
			frontwheel: wheel{
				radius: 22,
				psi:    42,
			},
		},
	}

	// INTIALIZING EMBEDDED STRUCT WITH EXISTING STRUCT INSTANCE
	myTaxi_one := taxi{
		owner: "sou",
		car:   myCar,
	}

	fmt.Println(myCar)
	fmt.Println(myTaxi)
	fmt.Println(myTaxi_one)
}
