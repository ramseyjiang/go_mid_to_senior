package httpreqtest

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"testing"
)

// TestRealCallSuccess and TestRealCallFail send a real internet call, in the test, it is bad thing.
// TestRealCallSuccess uses 0.54s to pass, TestRealCallFail uses 0.04s to pass. They are slower than mock requests.
// The right way, it is using mock not sending a request.
func TestRealCallSuccess(t *testing.T) {
	result, err := GetRepos("ramseyjiang")
	if err != nil {
		t.Error("TestRealCallSuccess failed.")
		return
	}
	if len(result) == 0 {
		t.Error("TestRealCallSuccess failed, array was empty.")
		return
	}
	if result[0]["full_name"] != "ramseyjiang/tools" {
		t.Error("TestRealCallSuccess failed, array was not sorted correctly.")
		return
	}
}

func TestRealCallFail(t *testing.T) {
	_, err := GetRepos("ramseyjiang")
	if err != nil {
		t.Error("TestRealCallFail failed.")
		return
	}
}

// Custom type that allows setting the func that our Mock Do func will run instead
// create a new type called MockDoType which has the same signature as the Go http package’s Do function.
type MockDoType func(req *http.Request) (*http.Response, error)

// MockClient is the mock client
// create a struct called MockClient which has a single property called MockDo which is our newly created MockDoType.
type MockClient struct {
	MockDo MockDoType
}

// Overriding what the Do function should "do" in our MockClient
// hang a Do method off our MockClient struct which also has the same signature as the http package’s Do function.
func (m *MockClient) Do(req *http.Request) (*http.Response, error) {
	return m.MockDo(req)
}

func TestFakeCallSuccess(t *testing.T) {
	// build our response JSON
	jsonResponse := `[{
	   "full_name": "mock-repo"
	  }]`
	// create a new reader with that JSON
	r := ioutil.NopCloser(bytes.NewReader([]byte(jsonResponse)))
	Client = &MockClient{
		MockDo: func(*http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       r,
			}, nil
		},
	}
	result, err := GetRepos("ramseyjiang")
	if err != nil {
		t.Error("TestFakeCallSuccess failed.")
		return
	}
	if len(result) == 0 {
		t.Error("TestFakeCallSuccess failed, array was empty.")
		return
	}
	if result[0]["full_name"] != "mock-repo" {
		t.Error("TestFakeCallSuccess failed, array was not sorted correctly.")
		return
	}
}
func TestFakeCallFail(t *testing.T) {
	// create a client that throws and returns an error
	Client = &MockClient{
		MockDo: func(*http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: http.StatusNotFound,
				Body:       nil,
			}, errors.New("mock Error")
		},
	}
	_, err := GetRepos("ramseyjiang")
	if err == nil {
		t.Error("TestFakeCallFail failed.")
		return
	}
}
