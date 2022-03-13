package restymock

import (
	"errors"
	"io"
	"net/http"
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func Test_GetXML(t *testing.T) {
	tests := map[string]struct {
		data     string
		path     string
		respCode int
		want     *XML
		wantErr  error
	}{
		"happy path": {
			data:     `<root><Name>Oleg</Name></root>`,
			path:     "https://example.com/mynameis",
			respCode: 200,
			want:     &XML{Name: "Oleg"},
			wantErr:  nil,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			defer httpmock.DeactivateAndReset()
			rst := resty.New()
			s := &Service{
				Client: rst,
			}

			httpmock.Activate()
			httpmock.ActivateNonDefault(rst.GetClient())
			httpmock.RegisterResponder("GET", tt.path, newResponder(tt.respCode, tt.data, "application/xml"))
			got, err := s.GetXML(tt.path)
			if tt.wantErr != nil && !errors.Is(err, tt.wantErr) {
				t.Errorf("GetXML() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_XMLError(t *testing.T) {
	tests := map[string]struct {
		data     string
		path     string
		respCode int
		want     *XML
		wantErr  bool
	}{
		"error response": {
			data:     ``,
			path:     "https://example.com/mynameis",
			respCode: 500,
			want:     nil,
			wantErr:  true,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			defer httpmock.DeactivateAndReset()
			rst := resty.New()
			s := &Service{
				Client: rst,
			}

			httpmock.Activate()
			httpmock.ActivateNonDefault(rst.GetClient())
			httpmock.RegisterResponder("GET", tt.path, newResponder(tt.respCode, tt.data, "application/xml"))
			got, err := s.GetXML(tt.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetXML() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_JSON(t *testing.T) {
	tests := map[string]struct {
		data     string
		path     string
		respCode int
		want     *JSON
		wantErr  error
	}{
		"happy path": {
			data:     `{"name":"Oleg"}`,
			path:     "https://example.com/mynameis",
			respCode: 200,
			want:     &JSON{Name: "Oleg"},
			wantErr:  nil,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			defer httpmock.DeactivateAndReset()
			rst := resty.New()
			s := &Service{
				Client: rst,
			}

			httpmock.ActivateNonDefault(rst.GetClient())
			httpmock.RegisterResponder("GET", tt.path, newResponder(tt.respCode, tt.data, "application/json"))
			got, err := s.GetJSON(tt.path)
			if tt.wantErr != nil && !errors.Is(err, tt.wantErr) {
				t.Errorf("GetJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_JSONStringResponder(t *testing.T) {
	// this test is skipped since it's intended to fail
	t.SkipNow()
	tests := map[string]struct {
		data     string
		path     string
		respCode int
		want     *JSON
		wantErr  error
	}{
		"happy path": {
			data:     `{"name":"Oleg"}`,
			path:     "https://example.com/mynameis",
			respCode: 200,
			want:     &JSON{Name: "Oleg"},
			wantErr:  nil,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			defer httpmock.DeactivateAndReset()
			rst := resty.New()
			s := &Service{
				Client: rst,
			}

			httpmock.ActivateNonDefault(rst.GetClient())
			httpmock.RegisterResponder("GET", tt.path, httpmock.NewStringResponder(tt.respCode, tt.data))
			got, err := s.GetJSON(tt.path)
			if tt.wantErr != nil && !errors.Is(err, tt.wantErr) {
				t.Errorf("GetJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_JSONJsonResponder(t *testing.T) {
	tests := map[string]struct {
		data     string
		path     string
		respCode int
		want     *JSON
		wantErr  error
	}{
		"happy path": {
			data:     `{"name":"Oleg"}`,
			path:     "https://example.com/mynameis",
			respCode: 200,
			want:     &JSON{Name: "Oleg"},
			wantErr:  nil,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			defer httpmock.DeactivateAndReset()
			rst := resty.New()
			s := &Service{
				Client: rst,
			}

			httpmock.ActivateNonDefault(rst.GetClient())
			httpmock.RegisterResponder("GET", tt.path, httpmock.NewJsonResponderOrPanic(tt.respCode, &JSON{Name: "Oleg"}))
			got, err := s.GetJSON(tt.path)
			if tt.wantErr != nil && !errors.Is(err, tt.wantErr) {
				t.Errorf("GetJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_MininalCase(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", "https://example.com", httpmock.NewStringResponder(200, "resp string"))
	resp, _ := http.Get("https://example.com")
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(resp.Body)
	body, _ := io.ReadAll(resp.Body)
	assert.Equal(t, "resp string", string(body))
}

func Test_MininalResty(t *testing.T) {
	rst := resty.New()
	httpmock.ActivateNonDefault(rst.GetClient())
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", "https://example.com", httpmock.NewStringResponder(200, "resp string"))
	resp, _ := rst.R().Get("https://example.com")
	assert.Equal(t, "resp string", resp.String())
}

func Test_MininalCustomJSON(t *testing.T) {
	rst := resty.New()
	httpmock.ActivateNonDefault(rst.GetClient())
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder(
		"GET",
		"https://example.com",
		newResponder(200, `{"name":"Oleg"}`, "application/json"),
	)
	resp, _ := rst.R().
		SetResult(&JSON{}).
		Get("https://example.com")

	assert.Equal(t, &JSON{Name: "Oleg"}, resp.Result().(*JSON))
}

func Test_MininalCustomXML(t *testing.T) {
	rst := resty.New()
	httpmock.ActivateNonDefault(rst.GetClient())
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder(
		"GET",
		"https://example.com",
		newResponder(
			200,
			`<root><Name>Oleg</Name></root>`,
			"application/xml",
		),
	)
	resp, _ := rst.R().
		SetResult(&XML{}).
		Get("https://example.com")

	assert.Equal(t, &XML{Name: "Oleg"}, resp.Result().(*XML))
}

func newResponder(s int, c string, ct string) httpmock.Responder {
	resp := httpmock.NewStringResponse(s, c)
	resp.Header.Set("Content-Type", ct)

	return httpmock.ResponderFromResponse(resp)
}
