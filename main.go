package main

import "log"
import "flag"
import "os"
import "os/signal"
import "syscall"

import . "github.com/caser789/go-rpc-framework/server"
import . "github.com/caser789/go-rpc-framework/client"

var port = flag.Uint("port", 1337, "port to listen  or connect to for rpc calls")
var isServer = flag.Bool("server", false, "activates server mode")
var http = flag.Bool("http", false, "whether it should use http")
var json = flag.Bool("json", false, "whether it should use json-rpc")
var serverSleep = flag.Duration("server.sleep", 0, "time for the server to sleep on requests")

func must(err error) {
	if err == nil {
		return
	}

	log.Panicln(err)
}

func runServer() {
    server := &Server{
        Port: *port,
        UseHttp: *http,
        UseJson: *json,
        Sleep: *serverSleep,
    }
    defer server.Close()

    go func() {
        handleSignals()
        server.Stop()
        os.Exit(0)
    }()

    must(server.Start())
    return
}

func runClient() {
	client := &Client{
		Port: *port,
        UseHttp: *http,
        UseJson: *json,
	}
    defer client.Close()

    must(client.Init())

	response, err := client.Execute("jiao")
	must(err)

	log.Println(response)
}

func handleSignals() {
	signals := make(chan os.Signal, 1)

	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	<-signals
	log.Println("signal received")
}

func main() {
	flag.Parse()

    if *http {
        log.Println("will use http")
    }
    if *json {
        log.Println("will use json-rpc")
    }

	if *isServer {
		log.Println("starting server")
		log.Printf("will listen on port %d\n", *port)
        runServer()
	}

	log.Println("starting client")
	log.Printf("will connect to port %d\n", *port)
    runClient()
    return
}
