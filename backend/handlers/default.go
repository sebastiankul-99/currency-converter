package handlers

import (
	"log"
	"net/http"
)

type DefaultHandler struct {
	l *log.Logger
}

func GetDefaultHandler(l *log.Logger) *DefaultHandler {
	return &DefaultHandler{l}
}
func (c *DefaultHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

	//rejecting calls that are not /currencies or /rate/
	http.Error(rw, "Page Not Found", http.StatusNotFound)

}
