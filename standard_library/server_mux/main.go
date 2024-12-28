package main


import (
	"net/http"
	"time"
)

func main(){
	person := http.NewServeMux()
	person.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request){
		w.Write([]byte("greetings!\n"))
	})

	dog := http.NewServeMux()
	dog.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request){
		w.Write([]byte("good puppy!\n"))
	})

	mux := http.NewServeMux()
	mux.Handle("/person/", http.StripPrefix("person", person))
	mux.Handle("/dog/", http.StripPrefix("/dog", dog))

	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request){
		w.Write([]byte("Hello!\n"))
	})

	s := http.Server{
		Addr: ":8000",
		ReadTimeout: 30 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout: 60 * time.Second,
		Handler: mux,
	}
	err := s.ListenAndServe()
	if err != nil {
		if err != http.ErrServerClosed{
			panic(err)
		}
	}

}