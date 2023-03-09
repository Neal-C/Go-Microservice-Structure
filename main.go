package main

import (
	"log"
)

func main() {
	service := NewCatFactService("https://catfact.ninja/fact")
	service = NewLoggingService(service)

	apiServer := NewAPIServer(service);

	log.Fatal(apiServer.start(":3333"));
}