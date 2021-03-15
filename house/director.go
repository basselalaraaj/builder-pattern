package house

type Builder interface {
	setHouseType() Builder
	setDoorType(string) Builder
	setWindowType(string) Builder
	setFloor(int) Builder
	getHouse() house
}

type director struct {
	builder Builder
}

type Director interface {
	SetBuilder(Builder)
	BuildHouse() house
}

func NewDirector() Director {
	return &director{}
}

func (d *director) SetBuilder(builder Builder) {
	d.builder = builder
}

func (d *director) BuildHouse() house {
	d.builder.setDoorType("6 panel")
	d.builder.setWindowType("slide")
	d.builder.setFloor(1)
	d.builder.setHouseType()
	return d.builder.getHouse()
}
