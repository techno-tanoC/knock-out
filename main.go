package main

import (
	"context"
	"fmt"
	"net/http"
	"os/exec"
	"time"
)

func main() {
	channel := make(chan struct{})

	m := http.NewServeMux()
	m.HandleFunc("/shutdown", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("finish server!!!")
		close(channel)
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
		exec.Command("ls", "-al").Run()
		fmt.Println("finish command!!!")
		close(channel)
	}()

	go func() {
		time.Sleep(time.Second * 1)
		fmt.Println("finish sleep!!!")
		close(channel)
	}()

	<-channel
	s.Shutdown(context.Background())
}
