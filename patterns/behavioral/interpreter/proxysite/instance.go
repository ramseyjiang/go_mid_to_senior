package proxysite

import (
	"crypto/tls"
	"net/http"
)

const secret = ""

// signatureRoundTripper the signature RoundTripper concrete struct.
type signatureRoundTripper struct {
	TLSClientConfig *tls.Config
	transport       *http.Transport
}

// RoundTrip Implements the Golang RoundTripper interface
func (t *signatureRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	newRequest := r.Clone(r.Context())
	signature, err := CreateRequestWithSignature(newRequest, "secrect_key_here")
	if err != nil {
		return nil, err
	}
	newRequest.Header.Set("X-Ca-Signature", signature)
	newRequest.Header.Set("X-Ca-Key", secret)
	return t.transport.RoundTrip(newRequest)
}

func CreateRequestWithSignature(req *http.Request, secretKey string) (signature string, err error) {
	return "", nil
}

type HikVisionAdapter struct {
	client *http.Client
}

func NewHikVisionAdapter() *HikVisionAdapter {
	return &HikVisionAdapter{
		client: &http.Client{
			Transport: &signatureRoundTripper{},
		},
	}
}
