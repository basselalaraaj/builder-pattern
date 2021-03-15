package house

type woodenHouse struct {
	houseType  string
	doorType   string
	windowType string
	floor      int
}

func WoodenHouseBuilder() *woodenHouse {
	return &woodenHouse{}
}

func (wh *woodenHouse) setHouseType() Builder {
	wh.houseType = "wood"
	return wh
}

func (wh *woodenHouse) setDoorType(doorType string) Builder {
	wh.doorType = doorType
	return wh
}
func (wh *woodenHouse) setWindowType(windowType string) Builder {
	wh.windowType = windowType
	return wh
}
func (wh *woodenHouse) setFloor(foor int) Builder {
	wh.floor = foor
	return wh
}
func (wh *woodenHouse) getHouse() house {
	return house{
		HouseType:  wh.houseType,
		doorType:   wh.doorType,
		windowType: wh.windowType,
		floor:      wh.floor,
	}
}
