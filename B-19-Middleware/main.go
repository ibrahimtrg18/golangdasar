package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const USERNAME = "batman"
const PASSWORD = "secret"

func MiddlewareAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()

		if !ok {
			w.Write([]byte("something is wrong!"))
			return
		}

		isValid := (username == USERNAME) && (password == PASSWORD)
		if isValid {
			w.Write([]byte("Wrong username/password"))
			return
		}

		next.ServeHTTP(w, r)
	})
}

func MiddlewareAllowOnlyGet(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			w.Write([]byte("Only GET is allowed"))
			return
		}
		next.ServeHTTP(w, r)
	})
}

func ActionStudent(w http.ResponseWriter, r *http.Request) {
	if id := r.URL.Query().Get("id"); id != "" {
		OutputJSON(w, SelectStudent(id))
		return
	}

	OutputJSON(w, GetStudent())
}

func OutputJSON(w http.ResponseWriter, o interface{}) {
	res, err := json.Marshal(o)

	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func main() {
	mux := http.DefaultServeMux

	mux.HandleFunc("/student", ActionStudent)

	var handler http.Handler = mux
	handler = MiddlewareAuth(handler)
	handler = MiddlewareAllowOnlyGet(handler)

	server := new(http.Server)
	server.Addr = ":9000"
	server.Handler = handler

	fmt.Println("Server started at localhost:9000")
	server.ListenAndServe()
}
