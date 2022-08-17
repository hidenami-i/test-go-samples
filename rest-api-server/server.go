package main

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net"
	"net/http"
	"time"
)

type UserEntity struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.RealIP)
	router.Use(middleware.RequestID)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Timeout(60 * time.Second))

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		MyRequestHandler(w, r)
	})

	srv := &http.Server{
		Addr:              net.JoinHostPort("localhost", "3000"),
		Handler:           router,
		TLSConfig:         nil,
		ReadTimeout:       0,
		ReadHeaderTimeout: 0,
		WriteTimeout:      0,
		IdleTimeout:       0,
		MaxHeaderBytes:    0,
		TLSNextProto:      nil,
		ConnState:         nil,
		ErrorLog:          nil,
		BaseContext:       nil,
		ConnContext:       nil,
	}

	err := srv.ListenAndServe()
	if err != nil {
		return
	}
}

func MyRequestHandler(w http.ResponseWriter, r *http.Request) {
	// here we read from the request context and fetch out `"user"` key set in
	// the MyMiddleware example above.
	//user := r.Context().Value("user").(string)

	// respond to the client
	//w.Write([]byte(fmt.Sprintf("hi %s", user)))
	entity := UserEntity{
		Id:   "sample",
		Name: "Tanaka Taro",
		Age:  20,
	}
	marshal, err := json.Marshal(entity)
	if err != nil {
		return
	}
	w.Write(marshal)
	//w.Write([]byte("welcome"))
}
