package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/sebastiankul-99/currency-converter/handlers"
)

func main() {
	l := log.New(os.Stdout, "currency-api", log.LstdFlags)
	hd := handlers.GetDefaultHandler(l)
	hc := handlers.GetCurrencytHandler(l)
	hr := handlers.GetRateHandler(l)
	sm := http.NewServeMux()

	sm.Handle("/currencies", hc)
	sm.Handle("/rate/", hr)
	sm.Handle("/", hd)
	s := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		IdleTimeout:  80 * time.Second,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}

	}()

	signalChan := make(chan os.Signal)
	signal.Notify(signalChan, os.Interrupt)
	signal.Notify(signalChan, os.Kill)
	sig := <-signalChan
	l.Println("Received termination, shutting down... ", sig)
	tc, err := context.WithTimeout(context.Background(), 40*time.Second)
	if err != nil {
		l.Println(err)
	}
	s.Shutdown(tc)
}
