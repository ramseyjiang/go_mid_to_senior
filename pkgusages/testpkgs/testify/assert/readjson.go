package assert

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

type Player struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func trigger() {
	_ = processData("data.txt")
}

func processData(file string) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer func(f *os.File) {
		err = f.Close()
		if err != nil {
			log.Println(err.Error())
		}
	}(f)
	return unmarshalAndPrint(f)
}

// unmarshalAndPrint is used to read json.
// For testing, instead of preparing data and opening a file, we just pass a literal JSON string to strings.NewReader
func unmarshalAndPrint(f io.Reader) error {
	data, err := io.ReadAll(f)
	if err != nil {
		return err
	}
	var players []Player

	err = json.Unmarshal(data, &players)
	if err != nil {
		return err
	}

	for _, p := range players {
		fmt.Println("Player name: ", p.Name)
	}
	return nil
}
