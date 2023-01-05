package main

import (
	"net/http"
	"strings"

	opamiddleware "github.com/Joffref/opa-middleware"
	"github.com/Joffref/opa-middleware/config"
)

var policy = `
package policy
default allow = false
document := {
    "users": {
        "alice": {
            "permissions": [
                "articles:read"
            ]
        },
        "bob": {
            "permissions": [
                "articles:read",
                "articles:write"
            ]
        }
    }
}
allow {
    input.method = "GET"
    input.path = ["articles"]
    has_permission("articles:read")
}
allow {
    input.method = "POST"
    input.path = ["articles"]
    has_permission("articles:write")
}
allow {
    input.method = "PUT"
    input.path = ["articles"]
    has_permission("articles:write")
}
allow {
    input.method = "DELETE"
    input.path = ["articles"]
    has_permission("articles:write")
}
has_permission(permission) {
    user := input.user
    permissions := document.users[user].permissions
    permissions[_] = permission
}
`

type Handler struct {
}

func (h *Handler) GetArticle(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get article"))
}

func (h *Handler) CreateArticle(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create article"))
}

func (h *Handler) UpdateArticle(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Update article"))
}

func (h *Handler) DeleteArticle(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete article"))
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.GetArticle(w, r)
	case http.MethodPost:
		h.CreateArticle(w, r)
	case http.MethodPut:
		h.UpdateArticle(w, r)
	case http.MethodDelete:
		h.DeleteArticle(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func main() {
	handler, err := opamiddleware.NewHTTPMiddleware(
		&config.Config{
			Policy: policy,
			Query:  "data.policy.allow",
			InputCreationMethod: func(r *http.Request) (map[string]interface{}, error) {
				return map[string]interface{}{
					"path":   strings.Split(r.URL.Path, "/")[1:],
					"method": r.Method,
					"user":   r.Header.Get("X-User"),
				}, nil
			},
			ExceptedResult:   true,
			DeniedStatusCode: http.StatusForbidden,
			DeniedMessage:    "Forbidden",
		},
		&Handler{},
	)
	if err != nil {
		panic(err)
	}
	http.HandleFunc("/articles", handler.ServeHTTP)
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
