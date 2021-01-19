package handlers

import "net/http"

type Hello struct {
}

func (h *Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
}
