package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"log"
	"net/http"
)

type Server struct {
	ID string
	Storage storage
}

type storage interface {
	write(datum *Datum) (Receipt, err error)
}

type arrayStorage struct {
	currentIndex int
	items []Datum
}

func main() {
	s := Server{
		ID: "TODO",
	}

	router := s.Routes()
	log.Fatalln(http.ListenAndServe(":7000", router))
}

func (s *Server) Routes() *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Get("/healthz", s.Healthz())

	router.Post("/data", s.CreateDataPoint())

	return router
}

func (s* arrayStorage) write(datum *Datum) (Receipt, error) {
	r := Receipt{}

	return r, nil
}

func (s *Server) CreateDataPoint() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (s *Server) Healthz() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if _, err := w.Write([]byte("OK!")); err != nil {
			fmt.Println("NOT OK")
		}
	}
}

