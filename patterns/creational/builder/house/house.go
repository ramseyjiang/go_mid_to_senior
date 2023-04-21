package house

type House struct {
	windowType string
	doorType   string
	floor      int
}

type Builder interface {
	setWindowType()
	setDoorType()
	setNumFloor()
	getHouse() House
}
