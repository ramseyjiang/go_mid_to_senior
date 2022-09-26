package client

import (
	"errors"
	"io"
	"net/http"
)

func Request() (err error) {
	resp, err := http.Get("http://localhost:8080/ping") // call the upstream service on the client side.
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
	}(resp.Body)
	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		return errors.New("bad response")
	}

	return nil
}
