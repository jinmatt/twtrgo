package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jinmatt/twtrgo/config"
	"github.com/jinmatt/twtrgo/http"
	"github.com/jinmatt/twtrgo/http/handler"
	"github.com/jinmatt/twtrgo/services"
)

func main() {
	log.Println("Server starting up...")

	config, err := config.NewConfig()
	if err != nil {
		log.Println("Error loading config:")
		log.Fatal(err.Error())
	}

	err = services.Init(config)
	if err != nil {
		log.Println("Error initializing services:")
		log.Fatal(err.Error())
	}

	handler := handler.NewHandler()
	server := http.NewServer(handler)

	err = server.Start(config)
	if err != nil {
		log.Fatalln("Unable to start server:", err)
	}

	// set up signal handling
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	// wait until we get a signal
	<-sig

	// stop handling signals
	signal.Stop(sig)
	signal.Reset()

	log.Println("Caught signal. Shutting down...")

	server.Stop(3 * time.Second)

	log.Println("HTTP server shutdown complete.")
}
