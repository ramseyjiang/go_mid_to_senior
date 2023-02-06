package server

import "net/http"

type server interface {
	handleRequest(string, string) (int, string)
}

type Web struct {
	application       *Application
	maxAllowedRequest int
	rateLimiter       map[string]int
}

func newWebServer() *Web {
	return &Web{
		application:       &Application{},
		maxAllowedRequest: 2,
		rateLimiter:       make(map[string]int),
	}
}

func (w *Web) handleRequest(url, method string) (int, string) {
	allowed := w.checkRateLimiting(url)
	if !allowed {
		return http.StatusForbidden, "Not Allowed"
	}
	return w.application.handleRequest(url, method)
}

func (w *Web) checkRateLimiting(url string) bool {
	if w.rateLimiter[url] == 0 {
		w.rateLimiter[url] = 1
	}
	if w.rateLimiter[url] > w.maxAllowedRequest {
		return false
	}
	w.rateLimiter[url] = w.rateLimiter[url] + 1
	return true
}

type Application struct {
}

func (a *Application) handleRequest(url, method string) (int, string) {
	if url == "/app/status" && method == "GET" {
		return http.StatusOK, "Ok"
	}

	if url == "/create/user" && method == "POST" {
		return http.StatusCreated, "User Created"
	}
	return http.StatusNotFound, "Not Ok"
}
