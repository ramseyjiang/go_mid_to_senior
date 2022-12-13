package db

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

type (
	animal struct {
		Name  string `json:"name"`
		Sound string `json:"sound"`
	}
)

var animals []animal

func init() {
	wd, err := os.Getwd()
	if err != nil {
		log.Println(err.Error())
	}

	file, err := os.Open(fmt.Sprintf("%v/db/data.json", wd))
	if err != nil {
		log.Println(err.Error())
	}

	byteFile, _ := io.ReadAll(file)
	if err != nil {
		log.Println(err.Error())
	}

	err = json.Unmarshal(byteFile, &animals)
	if err != nil {
		log.Println(err.Error())
	}
}

func GetAnimal(name string) (*animal, error) {
	log.Println(animals)
	for _, v := range animals {
		// log.Println(v.Name, v.Name == name)
		if name == v.Name {
			return &v, nil
		}
	}
	return nil, errors.New("no animal found")
}
