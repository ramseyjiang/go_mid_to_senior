package main

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// run "go test -v"
/**
go test -v -run <Test Function Name>

go test -v -run TestCheckHealth
=== RUN   TestCheckHealth
=== RUN   TestCheckHealth/Check_health_status
--- PASS: TestCheckHealth (0.00s)
    --- PASS: TestCheckHealth/Check_health_status (0.00s)
PASS
ok      github.com/ramseyjiang/go_mid_to_senior/projects/restbooks 0.011s

go test -v -run TestGetEntryByID
=== RUN   TestGetEntryByID
--- PASS: TestGetEntryByID (0.00s)
=== RUN   TestGetEntryByIDNotFound
--- PASS: TestGetEntryByIDNotFound (0.00s)
PASS
ok      github.com/ramseyjiang/go_mid_to_senior/projects/restbooks 0.014s
*/
func TestCheckHealth(t *testing.T) {
	t.Run("Check health status", func(t *testing.T) {
		req := httptest.NewRequest("GET", "https://mysite.com/example", nil)
		writer := httptest.NewRecorder()
		CheckHealth(writer, req)
		resp := writer.Result()
		body, _ := io.ReadAll(resp.Body)
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				log.Fatal(err)
			}
		}(resp.Body)

		assert.Equal(t, "health check passed", string(body))
	})
}
func TestGetBooks(t *testing.T) {
	reqBody := bytes.NewReader([]byte{})
	req, err := http.NewRequest("GET", "/books", reqBody)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetBooks)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `[{"id":2,"first_name":"Krish","last_name":"Bhanushali","email":"krishsb2405@gmail.com","phone_number":"7798775575"},{"id":3,"first_name":"Kelly","last_name":"Franco","email":"kelly.franco@gmail.com","phone_number":"1112223333"},{"id":4,"first_name":"John","last_name":"Doe","email":"john.doe@gmail.com","phone_number":"1234567890"},{"id":5,"first_name":"xyz","last_name":"pqr","email":"xyz@pqr.com","phone_number":"1234567890"},{"id":6,"first_name":"xyz","last_name":"pqr","email":"xyz@pqr.com","phone_number":"1234567890"},{"id":7,"first_name":"xyz","last_name":"pqr","email":"xyz@pqr.com","phone_number":"1234567890"},{"id":8,"first_name":"xyz","last_name":"pqr","email":"xyz@pqr.com","phone_number":"1234567890"},{"id":9,"first_name":"xyz","last_name":"pqr","email":"xyz@pqr.com","phone_number":"1234567890"}]`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestGetBookByID(t *testing.T) {
	reqBody := bytes.NewReader([]byte{})
	req, err := http.NewRequest("GET", "/book", reqBody)
	if err != nil {
		t.Fatal(err)
	}

	q := req.URL.Query()
	q.Add("id", "2")
	req.URL.RawQuery = q.Encode()
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetBookByID)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"id":2,"first_name":"Krish","last_name":"Bhanushali","email":"krishsb2405@gmail.com","phone_number":"7798775575"}`

	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestGetBookByIDNotFound(t *testing.T) {
	reqBody := bytes.NewReader([]byte{})
	req, err := http.NewRequest("GET", "/book", reqBody)

	if err != nil {
		t.Fatal(err)
	}

	q := req.URL.Query()
	q.Add("id", "123")
	req.URL.RawQuery = q.Encode()
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetBookByID)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status == http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}

func TestCreateBook(t *testing.T) {
	var jsonStr = []byte(`{"first_name":"xyz","last_name":"pqr","email":"xyz@pqr.com","phone_number":"1234567890"}`)

	req, err := http.NewRequest("POST", "/book", bytes.NewBuffer(jsonStr))

	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateBook)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `{"id":16,"first_name":"xyz","last_name":"pqr","email":"xyz@pqr.com","phone_number":"1234567890"}`

	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestEditBook(t *testing.T) {
	var jsonStr = []byte(`{"id":15,"first_name":"xyz change","last_name":"pqr","email":"xyz@pqr.com","phone_number":"1234567890"}`)

	req, err := http.NewRequest("PUT", "/book", bytes.NewBuffer(jsonStr))

	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(UpdateBook)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `{"id":15,"first_name":"xyz change","last_name":"pqr","email":"xyz@pqr.com","phone_number":"1234567890"}`

	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
func TestDeleteEntry(t *testing.T) {
	reqBody := bytes.NewReader([]byte{})
	req, err := http.NewRequest("DELETE", "/book", reqBody)

	if err != nil {
		t.Fatal(err)
	}

	q := req.URL.Query()
	q.Add("id", "15")
	req.URL.RawQuery = q.Encode()
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(DeleteBook)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `{"id":15,"first_name":"xyz change","last_name":"pqr","email":"xyz@pqr.com","phone_number":"1234567890"}`

	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
