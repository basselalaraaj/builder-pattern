package house

type iglooHouse struct {
	houseType  string
	doorType   string
	windowType string
	floor      int
}

func IglooHouseBuilder() *iglooHouse {
	return &iglooHouse{}
}

func (ih *iglooHouse) setHouseType() Builder {
	ih.houseType = "igloo"
	return ih
}

func (ih *iglooHouse) setDoorType(doorType string) Builder {
	ih.doorType = doorType
	return ih
}
func (ih *iglooHouse) setWindowType(windowType string) Builder {
	ih.windowType = windowType
	return ih
}
func (ih *iglooHouse) setFloor(foor int) Builder {
	ih.floor = foor
	return ih
}
func (ih *iglooHouse) getHouse() house {
	return house{
		HouseType:  ih.houseType,
		doorType:   ih.doorType,
		windowType: ih.windowType,
		floor:      ih.floor,
	}
}
