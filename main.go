package main

import "log"
import "flag"
import "os"
import "os/signal"
import "syscall"

import "github.com/caser789/go-rpc-framework/server"
import "github.com/caser789/go-rpc-framework/client"

var port = flag.Uint("port", 1337, "port to listen  or connect to for rpc calls")
var isServer = flag.Bool("server", false, "activates server mode")

func must(err error) {
	if err == nil {
		return
	}

	log.Panicln(err)
}

func handleSignals() {
	signals := make(chan os.Signal, 1)

	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	<-signals
	log.Println("signal received")
}

func main() {
	flag.Parse()

	if *isServer {
		log.Println("starting server")
		log.Printf("will listen on port %d\n", *port)

		server := &Server{
			Port: *port,
		}

		go func() {
			handleSignals()
			server.Stop()
			os.Exit(0)
		}()

		must(server.Start())
		return
	}

	log.Println("starting client")
	log.Printf("will connect to port %d\n", *port)

	client := &Client{
		Port: *port,
	}

	response, err := client.Execute("jiao")
	must(err)

	log.Println(response)
}
