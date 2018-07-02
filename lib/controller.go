package lib

import (
	"io/ioutil"
	"fmt"
	"net/http"
	"log"
	"time"
)

type Middleware func(handlerFunc http.HandlerFunc) http.HandlerFunc

func Method(m string)Middleware{
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			if r.Method != m {
				http.Error(w,"Bad Request",http.StatusBadRequest)
				return
			}
			f.ServeHTTP(w,r)
		}
	}
}

func Logger() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(writer http.ResponseWriter, request *http.Request) {
			start := time.Now()
			defer func() { log.Printf("\n%v-%v\n", time.Since(start), request.RequestURI) }()
			f.ServeHTTP(writer, request)
		}
	}
}

func GetHooksfunc(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		fmt.Fprintln(w, "Bad Request")
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(string(body))
}

func Chain(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		f = m(f)
	}
	return f
}
