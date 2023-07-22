package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/ping", func(rw http.ResponseWriter, r *http.Request) {
		_, _ = rw.Write([]byte("pong"))
	})

	ch := handlers.CORS(handlers.AllowedOrigins([]string{"*"}))

	port := "9000"
	s := http.Server{
		Addr:         "localhost:" + port,
		Handler:      ch(r),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	log.Println("Server started at port ", port)

	err := s.ListenAndServe()
	if err != nil {
		log.Fatalln(err)
	}

}
