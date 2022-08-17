package main

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

type UserEntity struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		MyRequestHandler(w, r)
	})
	err := http.ListenAndServe(":3000", router)
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
