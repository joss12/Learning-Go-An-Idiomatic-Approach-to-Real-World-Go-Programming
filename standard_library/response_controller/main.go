package main

import(
	"fmt"
	"errors"
	"log/slog"
	"net/http"
	"time"
)

func doStuff(i int)string{
	time.Sleep(1 * time.Second)
	return fmt.Sprintf("%d\n", i *2)
}

func handler(rw http.ResponseWriter, req *http.Request){
	rc := http.NewResponseController(rw)
	for i := 0; i < 10; i++{
		result := doStuff(i)
		_, err := rw.Write([]byte(result))
		if err != nil{
			slog.Error("error write", "msg", err)
			return
		}
		err = rc.Flush()
		if err != nil && !errors.Is(err, http.ErrNotSupported){
			slog.Error("error flushing", "msg", err)
			return
		}
	}
}

func main(){
	s := http.Server{
		Addr:         ":8080",
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 6 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      http.HandlerFunc(handler),
	}
	fmt.Println("Server Started...")
	err := s.ListenAndServe()
	if  err != nil{
		if !errors.Is(err, http.ErrServerClosed){
			panic(err)
		}
	}
}
