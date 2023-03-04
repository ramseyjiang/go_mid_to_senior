package purse

import (
	"github.com/pkg/errors"
)

// SecurityCode is the complex subsystem's structs.
type SecurityCode struct {
	code int
}

func NewSecurityCode(code int) *SecurityCode {
	return &SecurityCode{
		code: code,
	}
}

func (s *SecurityCode) checkCode(incomingCode int) error {
	if s.code != incomingCode {
		return errors.New("security Code is incorrect")
	}

	return nil
}
