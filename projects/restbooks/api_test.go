package main

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
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
	expected := `[{"id":"079deccd-514f-4be2-9904-ea6482198e29","first_name":"First","last_name":"test","email":"first@gmail.com","mobile":4723913},{"id":"269e72f5-b6cf-4afc-bf54-0b7060ce6b8f","first_name":"Second","last_name":"test","email":"second@gmail.com","mobile":4723914}]`
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
	q.Add("id", "269e72f5-b6cf-4afc-bf54-0b7060ce6b8f")
	req.URL.RawQuery = q.Encode()
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetBookByID)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"id":"269e72f5-b6cf-4afc-bf54-0b7060ce6b8f","first_name":"Second","last_name":"test","email":"second@gmail.com","mobile":4723914}`

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
	var jsonStr = []byte(`{"id":"` + uuid.New().String() + `","first_name":"xyz","last_name":"pqr","email":"xyz@pqr.com","mobile":123478}`)

	req, err := http.NewRequest("POST", "/book", bytes.NewBuffer(jsonStr))
	log.Println(bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateBook)
	handler.ServeHTTP(rr, req)
	log.Println(rr.Body)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestEditBook(t *testing.T) {
	var jsonStr = []byte(`{"id":"132cb11e-e0f0-44b9-9c62-e19c8dd0a178","first_name":"xyz change","last_name":"pqr","email":"xyz@pqr.com","mobile":12345690}`)

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

	expected := `{"id":"132cb11e-e0f0-44b9-9c62-e19c8dd0a178","first_name":"xyz change","last_name":"pqr","email":"xyz@pqr.com","mobile":12345690}`

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
	q.Add("id", "12aa9259-0fbd-45bd-8165-6f7d0ef11235")
	req.URL.RawQuery = q.Encode()
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(DeleteBook)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `{"error":"No book found with the id=12aa9259-0fbd-45bd-8165-6f7d0ef11235"}`

	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
