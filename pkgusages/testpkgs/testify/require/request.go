package require

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type PlayerInfo struct {
	Name     string
	Team     string
	Position string
}

func savePlayerInfo(playerInfo PlayerInfo, url string) error {
	if playerInfo.Name == "" || playerInfo.Position == "" || playerInfo.Team == "" {
		return fmt.Errorf("missing data")
	}
	b, err := json.Marshal(playerInfo)
	if err != nil {
		return err
	}
	body := bytes.NewBuffer(b)
	req, err := http.NewRequest("POST", url, body)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	_, err = client.Do(req)

	return err
}
