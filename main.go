package main

import (
	"fmt"
	"net/http"
	"strconv"
	"sync/atomic"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	var counter atomic.Uint64

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/counter", func(w http.ResponseWriter, r *http.Request) {
		answer := fmt.Sprintf("The counter is %d", counter.Load())
		w.Write([]byte(answer))
	})

	r.Post("/counter", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		new_value, err := strconv.Atoi(r.FormValue("value"))
		if err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			w.Write([]byte("Invalid value"))
			return
		}

		counter.Store(uint64(new_value))
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("Counter saved"))
	})

	http.ListenAndServe(":3000", r)
}

