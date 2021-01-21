package handlers

import (
	"fmt"
	"log"
	"net/http"
)

type Achievement struct {
	l *log.Logger
}

func NewAchievement(l *log.Logger) *Achievement {
	return &Achievement{l}
}

func (h *Achievement) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.l.Println("Achievement")
	data := 20
	fmt.Fprintf(w, "Level %d", data)
}
