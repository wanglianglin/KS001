package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/braintree/manners"
	"github.com/wanglianglin/ks001/app/handlers"
	"github.com/wanglianglin/ks001/app/user"
)

const version = "1.0.0"

func main() {
	var (
		httpAddr   = flag.String("http", "0.0.0.0:80", "HTTP service address.")
		secret     = flag.String("secret", "secret", "JWT signing secret.")
	)
	flag.Parse()

	log.Println("Starting Auth service...")
	log.Printf("HTTP service listening on %s", *httpAddr)

	errChan := make(chan error, 10)

	mux := http.NewServeMux()
	mux.Handle("/login", handlers.LoginHandler(*secret, user.DB))
	mux.Handle("/version", handlers.VersionHandler(version))

	httpServer := manners.NewServer()
	httpServer.Addr = *httpAddr
	httpServer.Handler = handlers.LoggingHandler(mux)

	go func() {
		errChan <- httpServer.ListenAndServe()
	}()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	for {
		select {
		case err := <-errChan:
			if err != nil {
				log.Fatal(err)
			}
		case s := <-signalChan:
			log.Println(fmt.Sprintf("Captured %v. Exiting...", s))
			httpServer.BlockingClose()
			os.Exit(0)
		}
	}
}
