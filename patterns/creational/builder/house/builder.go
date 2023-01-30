package house

const usual = "usual"
const igloo = "igloo"

type Builder interface {
	setWindowType()
	setDoorType()
	setNumFloor()
	getHouse() House
}

type House struct {
	windowType string
	doorType   string
	floor      int
}

func getBuilder(builderType string) Builder {
	if builderType == "usual" {
		return newUsualBuilder()
	}

	if builderType == "igloo" {
		return newIglooBuilder()
	}
	return nil
}
