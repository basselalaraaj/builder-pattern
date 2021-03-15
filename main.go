package main

import (
	"fmt"

	"github.com/basselalaraaj/builder-pattern/house"
)

func main() {
	director := house.NewDirector()

	// wooden house builder
	woodenHouseBuilder := house.WoodenHouseBuilder()
	director.SetBuilder(woodenHouseBuilder)

	woodenHouse := director.BuildHouse()
	fmt.Printf("House type is %s \n", woodenHouse.HouseType)

	// igloo house builder
	iglooHouseBuilder := house.IglooHouseBuilder()
	director.SetBuilder(iglooHouseBuilder)

	iglooHouse := director.BuildHouse()
	fmt.Printf("House type is %s \n", iglooHouse.HouseType)
}
