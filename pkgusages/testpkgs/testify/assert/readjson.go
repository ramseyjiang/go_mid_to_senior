package assert

import (
	"encoding/json"
	"fmt"
	"io"
)

type Player struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
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
		return fmt.Errorf("json format error")
	}

	for _, p := range players {
		fmt.Println("Player name: ", p.Name)
	}
	return nil
}
