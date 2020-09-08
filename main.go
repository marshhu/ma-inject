package main

import (
	"context"
	"github.com/marshhu/ma-inject/inject"
	"github.com/marshhu/ma-inject/router"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	Init()
	Run()
}

func Init() {
	inject.Init()
}

func Run() {
	router := router.Init()

	s := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    time.Duration(10) * time.Second,
		WriteTimeout:   time.Duration(10) * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	go func() {
		log.Println("Server Listen at:8080")
		if err := s.ListenAndServe(); err != nil {
			log.Printf("Listen:%s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	log.Println("Shutdown Server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
