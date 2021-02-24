package main

import (
	"context"
	"fmt"
	"net/http"
	"os/exec"
	"time"
)

func main() {
	channel := make(chan *struct{})

	m := http.NewServeMux()
	m.HandleFunc("/shutdown", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("finish server!!!")
		channel <- nil
	})
	m.HandleFunc("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
	})

	s := &http.Server{
		Addr:    ":8080",
		Handler: m,
	}

	go func() {
		s.ListenAndServe()
	}()

	go func() {
		exec.Command("sleep", "1").Run()
		fmt.Println("finish command!!!")
		channel <- nil
	}()

	go func() {
		time.Sleep(time.Second * 1)
		fmt.Println("finish sleep!!!")
		channel <- nil
	}()

	<-channel
	s.Shutdown(context.Background())
}
