package handlers

import (
	"log"
	"net/http"
)

type Shortner struct{}

func NewShortnerHandler() *Shortner {
	return &Shortner{}
}

func (s *Shortner) Handle(w http.ResponseWriter, r *http.Request) {
	m := r.Method
	switch m {
	case "POST":
		s.handleGet(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method Not Allowed"))
	}
}

func (s *Shortner) handleGet(w http.ResponseWriter, r *http.Request) {
	log.Println("HULULU")
	w.Write([]byte("HULULU"))
}

// TODO LINK APP CORE
// TODO ADD MIDDLEWARE
