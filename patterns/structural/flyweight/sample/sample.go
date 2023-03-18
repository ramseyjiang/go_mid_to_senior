package sample

import "fmt"

// Flyweight interface that contains the shared intrinsic data.
type Flyweight interface {
	Operation(extrinsicData interface{})
}

// ConcreteFlyweight is used to implement the concrete flyweight struct that contains the shared intrinsic data.
type ConcreteFlyweight struct {
	intrinsicData interface{}
}

func (c *ConcreteFlyweight) Operation(extrinsicData interface{}) {
	fmt.Printf("ConcreteFlyweight: %v, %v\n", c.intrinsicData, extrinsicData)
}

// FlyweightFactory struct is used to create and manage the flyweights.
type FlyweightFactory struct {
	flyweights map[interface{}]Flyweight
}

func (f *FlyweightFactory) GetFlyweight(key string) Flyweight {
	if flyweight, ok := f.flyweights[key]; ok {
		return flyweight
	}
	flyweight := &ConcreteFlyweight{key}
	f.flyweights[key] = flyweight
	return flyweight
}

// Client is the client is defined to use the flyweight objects.
type Client struct {
	factory *FlyweightFactory
}

func (c *Client) Run() {
	flyweight1 := c.factory.GetFlyweight("key1")
	flyweight1.Operation("extrinsicData1")

	flyweight2 := c.factory.GetFlyweight("key2")
	flyweight2.Operation("extrinsicData2")
}
