package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"
)

func main(){
	terribleSecurity := TerribleSecurityProvider("GOPHER")

	mux := http.NewServeMux()

	//to apply the middleware to just the single route
	mux.Handle("/hello", terribleSecurity(RequestTimer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
			w.Write([]byte("hello!/n"))
		}))))

		s := http.Server{
			Addr: ":8080",
			Handler: mux,
		}

		fmt.Println("Server started...")
		
		err := s.ListenAndServe()
		if err != nil{
			if err != http.ErrServerClosed{
				panic(err)
			}
		}
}

func RequestTimer(h http.Handler)http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		start := time.Now()
		h.ServeHTTP(w, r)
		end := time.Now()
		slog.Info("request time", "path", r.URL.Path, "duration", end.Sub(start))
	})
}

var securityMsg = []byte("You don't give the secret password\n")

func TerribleSecurityProvider(password string)func(http.Handler)http.Handler{
	return func(h http.Handler)http.Handler{
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
			if r.Header.Get("X-Secret-Password")!= password{
				w.WriteHeader(http.StatusUnauthorized)
				w.Write(securityMsg)
				return
			}
			h.ServeHTTP(w, r)
		})
	}
}