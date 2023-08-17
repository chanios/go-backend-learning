package main

import (
	"backend/db"
	"backend/handler"
	"backend/router"
	"backend/store"
	"fmt"
	"os"
	"os/signal"
)

func main() {
	r := router.New()

	d := db.New()
	db.AutoMigrate(d)

	bs := store.NewBookStore(d)
	h := handler.NewHandler(bs)
	h.Register(r)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	s := make(chan struct{})

	go func() {
		_ = <-c
		fmt.Println("Gracefully shutting down...")
		_ = r.Shutdown()
		s <- struct{}{}
	}()

	err := r.Listen("127.0.0.1:3000")

	if err != nil {
		fmt.Printf("%v", err)
	}

	<-s

}
