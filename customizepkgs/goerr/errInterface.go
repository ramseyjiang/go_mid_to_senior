package goerr

import (
	"fmt"
	"strings"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

// Error checking can be simplified by defining an explicit error interface definition with the implementation.
// The error can occur in any layer and it is necessary to provide the option for each layer to interpret
// the error and further wrap with additional contextual information without losing the original error value.
// The GoError can be further decorated with the Causes which will hold the entire error stack.
type Error interface {
	error

	Code() string
	Message() string
	Cause() error
	Causes() []error
	Data() map[string]interface{}
	String() string
	ResponseErrType() ResponseErrType
	SetResponseType(r ResponseErrType) Error
	Component() ErrComponent
	SetComponent(c ErrComponent) Error
	Retryable() bool // The retry component can decide whether to retry for the error by having the flag Retryable.
	SetRetryable() Error
}

type GoError struct {
	code         string
	message      string
	data         map[string]interface{}
	causes       []error
	component    ErrComponent
	responseType ResponseErrType
	retryable    bool
	appendCause  bool
}

// ErrComponent will help to identify the layer where the error has occurred, and unnecessary error wraps can be avoided.
// For example, if the error component of servicetype occurs in the service layer,
// then the wrapping error might not be required. Component information checks will help to prevent exposing
// the errors which a user shouldn’t be informed about, like Database errors
type ErrComponent string

const (
	ErrService ErrComponent = "service"
	ErrRepo    ErrComponent = "repository"
	ErrLib     ErrComponent = "library"
)

// ResponseErrType will support the error categorization for easy interpretation.
type ResponseErrType string

const (
	BadRequest    ResponseErrType = "BadRequest"
	Forbidden     ResponseErrType = "Forbidden"
	NotFound      ResponseErrType = "NotFound"
	AlreadyExists ResponseErrType = "AlreadyExists"
)

func NewGoError(code string, message string, data map[string]interface{}, cause error) Error {
	return &GoError{
		code:    code,
		message: message,
		data:    data,
		causes:  []error{cause},
	}
}

func (e *GoError) Error() string {
	s := fmt.Sprintf("%s:%s", e.code, e.Message())
	if e.appendCause {
		s += getCauses(e.causes)
	}
	return s
}

func (e *GoError) Code() string {
	return e.code
}

func (e *GoError) Message() string {
	return localize(e.code, templates[e.code], e.data)
}

func (e *GoError) Cause() error {
	if len(e.causes) > 1 {
		return e.causes[0]
	}
	return nil
}

func (e *GoError) Causes() []error {
	return e.causes
}

func (e *GoError) Data() map[string]interface{} {
	return e.data
}

func (e *GoError) String() string {
	return e.Error()
}

func (e *GoError) ResponseErrType() ResponseErrType {
	return e.responseType
}

func (e *GoError) SetResponseType(r ResponseErrType) Error {
	e.responseType = r
	return e
}

func (e *GoError) Component() ErrComponent {
	return e.component
}

func (e *GoError) SetComponent(c ErrComponent) Error {
	e.component = c
	return e
}
func (e *GoError) Retryable() bool {
	return e.retryable
}

func (e *GoError) SetRetryable() Error {
	e.retryable = true
	return e
}

// AppendCause will append the error cause to the err string
func (e *GoError) AppendCause() Error {
	e.appendCause = true
	return e
}

func getCauses(errors []error) string {
	var s strings.Builder
	for _, err := range errors {
		s.WriteString(err.Error())
		s.WriteString("; ")
	}
	return s.String()
}

// i18N bundle file can be used to templatizing the error string
var templates = map[string]string{
	"ResourceNotFound": "Resource of type {{.kind}} '{{.id}}' is not found",
}

func localize(id, template string, data map[string]interface{}) string {
	bundle := i18n.NewBundle(language.English)
	localized := i18n.NewLocalizer(bundle, "en")
	return localized.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    id,
			Other: template,
		},
		TemplateData: data,
	})
}
